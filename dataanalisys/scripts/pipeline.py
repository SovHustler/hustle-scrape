import random

import pandas as pd
import typing


def bootstrap_by_category_labels(data: pd.DataFrame, category_fraction: float, dfs_count: int, min_judge_crosses: int, random_state: int = 42) -> typing.List[pd.DataFrame]:
    unique_labels = pd.Series(data.category_label.unique())

    random.seed(random_state)

    result = []
    for i in range(10 * dfs_count):
        random_state = random.randint(0, 1000000)
        category_set = unique_labels.sample(frac=category_fraction, random_state=random_state)
        index = data.category_label.isin(category_set)

        df = data[index]

        min_df_judge_crosses = get_min_judge_crosses(df)
        if min_df_judge_crosses >= min_judge_crosses:
            result.append(df)
        else:
            print("skipping %d as min judge crosses is %d (below %d)" % (i, min_df_judge_crosses, min_judge_crosses))

        if len(result) >= dfs_count:
            break

    return result


def get_min_judge_crosses(data: pd.DataFrame) -> int:
    return data.judge_name.value_counts().min()


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
    crosses_data['category_label'] = make_category_labels(crosses_data)

    return crosses_data


def make_category_labels(crosses_data: pd.DataFrame) -> pd.Series:
    return crosses_data.competition_id + crosses_data.category


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
