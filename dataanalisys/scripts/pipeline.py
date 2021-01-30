import pandas as pd


def undersample_to_least_popular_judge(data: pd.DataFrame) -> pd.DataFrame:
    min_judge_records = data.judge_name.value_counts().min()
    judge_names = data.judge_name.unique()

    dfs = [data[data.judge_name == judge_name].sample(min_judge_records) for judge_name in judge_names]
    return pd.concat(dfs)


def remove_not_popular_competitor_records(data: pd.DataFrame, min_records: int) -> pd.DataFrame:
    df = data.join(get_competitor_cross_counts(data), on="competitor_id")
    df = df[df.cross_counts >= min_records]
    df.drop(columns=["cross_counts"], inplace=True)
    return df


def get_competitor_cross_counts(data: pd.DataFrame) -> pd.DataFrame:
    result = data.competitor_id.value_counts()
    result.name = "cross_counts"
    return result


def remove_not_popular_judge_records(data: pd.DataFrame, min_records: int) -> pd.DataFrame:
    df = data.join(get_judge_cross_counts(data), on="judge_name")
    df = df[df.judge_cross_counts >= min_records]
    df.drop(columns=["judge_cross_counts"], inplace=True)
    return df


def get_judge_cross_counts(data: pd.DataFrame) -> pd.DataFrame:
    result = data.judge_name.value_counts()
    result.name = "judge_cross_counts"
    return result


def join_crosses_and_results(crosses_data: pd.DataFrame, results_data: pd.DataFrame) -> pd.DataFrame:
    return crosses_data.merge(results_data, on=["competitor_id", "category", "competition_id"])


def enrich_crosses_df(crosses_data: pd.DataFrame) -> pd.DataFrame:
    crosses_data['syntethic_competitor_id'] = make_synthetic_competitor_ids(crosses_data.competitor_id)
    crosses_data['unique_id'] = make_unique_competition_id(crosses_data)

    return crosses_data


def make_unique_competition_id(crosses_data: pd.DataFrame) -> pd.Series:
    return crosses_data.competitor_id.astype(str) + crosses_data.competition_id + crosses_data.category


def make_synthetic_competitor_ids(competitor_ids: pd.Series) -> pd.Series:
    id_map = {}
    i = 0
    for id in competitor_ids:
        if id not in id_map:
            id_map[id] = i
            i += 1

    return pd.Series([id_map[id] for id in competitor_ids])


def get_raw_places(file_path: str) -> pd.DataFrame:
    df = pd.read_csv(
        file_path,
        names=["competitor_id", "category", "competition_id", "total_competitors", "place_low", "place_high"],
        dtype={
            "competitor_id": str,
        }
    ).dropna()  # only registered competitors are taken into account

    df = df[df.competitor_id != "дебют"]
    df.competitor_id = df.competitor_id.astype(int)
    df['place'] = (df.place_low + df.place_high) / 2.

    return df


def get_raw_crosses(file_path: str) -> pd.DataFrame:
    df = pd.read_csv(
        file_path,
        names=["competitor_id", "competition_id", "judge_name", "category", "phase", "passed"],
        dtype={
            "competitor_id": str,
        }
    ).dropna()  # only registered competitors are taken into account

    df.competitor_id = df.competitor_id.astype(int)

    df.passed[df.passed == 0] = -1
    return df
