import pandas as pd
import typing


def get_total_diff_based_model_data(data: pd.DataFrame) -> pd.DataFrame:
    return get_total_model_data(
        make_diff_based_model_x(data),
        get_places_data(data)
    )


def get_total_cross_based_model_data(data: pd.DataFrame) -> pd.DataFrame:
    return get_total_model_data(
        make_cross_based_model_x(data),
        get_places_data(data)
    )


def get_total_model_data(x_data: pd.DataFrame, places: pd.DataFrame) -> pd.DataFrame:
    return x_data \
        .merge(places, on="unique_id") \
        .drop(columns=["unique_id"])


def get_places_data(data: pd.DataFrame) -> pd.DataFrame:
    return data[["unique_id", "place"]] \
        .groupby("unique_id") \
        .first() \
        .reset_index()


class JudgeMapping:
    mapping: dict
    columns: list

    def __init__(self, mapping: dict, columns: list):
        self.mapping = mapping
        self.columns = columns


class RecordCreator:
    common_columns = ["unique_id", "judge_name", "total_competitors"]
    extra_columns: list
    get_key: typing.Callable[[dict], typing.Any]
    get_val: typing.Callable[[dict], float]

    def __init__(self, extra_columns: list, get_key: typing.Callable[[dict], tuple],
                 get_val: typing.Callable[[dict], float]):
        self.extra_columns = extra_columns
        self.get_key = get_key
        self.get_val = get_val


def make_diff_based_model_x(data: pd.DataFrame) -> pd.DataFrame:
    record_creator = RecordCreator(
        ["competition_score"],
        lambda d: d["judge_name"],
        lambda d: d["competition_score"]
    )

    source_data = get_diff_based_model_source_data(data)

    return make_model_x(
        source_data,
        make_diff_based_model_judge_mapping(source_data),
        record_creator
    )


def make_cross_based_model_x(data: pd.DataFrame) -> pd.DataFrame:
    record_creator = RecordCreator(
        ["passed", "cross_count"],
        lambda d: (d["judge_name"], d["passed"]),
        lambda d: d["cross_count"]
    )

    source_data = get_cross_based_model_source_data(data)

    return make_model_x(
        source_data,
        make_cross_based_model_judge_mapping(source_data),
        record_creator
    )


def make_model_x(source_data: pd.DataFrame, judge_mapping: JudgeMapping,
                 record_creator: RecordCreator) -> pd.DataFrame:
    all_records = []
    all_columns = record_creator.common_columns + record_creator.extra_columns

    def make_record() -> list:
        return [0] * (len(judge_mapping.mapping) + 2)

    def make_record_dict(tup: tuple) -> dict:
        result = {}
        for key, val in zip(all_columns, tup):
            result[key] = val

        return result

    curr_id = None
    curr_record = make_record()

    record_iterator = zip(*[source_data[col].values for col in all_columns])

    for tup in record_iterator:
        id, judge_name, total_competitors, *extra_fields = tup

        if curr_id is None:
            curr_id = id

        if id != curr_id:
            all_records.append(curr_record)
            curr_record = make_record()
            curr_id = id

        record_dict = make_record_dict(tup)

        curr_record[
            judge_mapping.mapping[
                record_creator.get_key(record_dict)
            ]
        ] = record_creator.get_val(record_dict)
        curr_record[len(judge_mapping.mapping)] = total_competitors
        curr_record[len(judge_mapping.mapping) + 1] = id

    return pd.DataFrame(all_records, columns=judge_mapping.columns + ["total_competitors", "unique_id"])


def make_diff_based_model_judge_mapping(data: pd.DataFrame) -> JudgeMapping:
    all_judges = data.judge_name.unique()

    judge_mapping = {}
    cnt = 0
    judge_mapping_columns = []
    for judge in all_judges:
        judge_mapping[judge] = cnt
        judge_mapping_columns.append(judge)
        cnt += 1

    return JudgeMapping(judge_mapping, judge_mapping_columns)


def make_cross_based_model_judge_mapping(data: pd.DataFrame) -> JudgeMapping:
    all_judges = data.judge_name.unique()

    judge_mapping = {}
    cnt = 0
    judge_mapping_columns = []
    for judge in all_judges:
        for passed in [-1, 1]:
            judge_mapping[(judge, passed)] = cnt
            judge_mapping_columns.append(judge + "_" + str(passed))
            cnt += 1

    return JudgeMapping(judge_mapping, judge_mapping_columns)


def get_diff_based_model_source_data(data: pd.DataFrame) -> pd.DataFrame:
    result = get_cross_based_model_source_data(data)
    result["competition_score"] = result.passed * result.cross_count
    return result[["unique_id", "judge_name", "total_competitors", "competition_score"]] \
        .groupby(["unique_id", "judge_name", "total_competitors"]) \
        .sum() \
        .reset_index()


def get_cross_based_model_source_data(data: pd.DataFrame) -> pd.DataFrame:
    data_columns = ["unique_id", "judge_name", "passed"]
    result = data[data_columns].groupby(data_columns).size().reset_index(name='cross_count')

    competitors = data[["unique_id", "total_competitors"]].groupby(["unique_id"]).first().reset_index()
    return result.merge(competitors, on="unique_id")
