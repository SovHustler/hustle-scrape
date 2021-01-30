import pandas as pd


def split_judge_crosses_df(data: pd.DataFrame) -> [pd.DataFrame, pd.DataFrame]:
    return data[data.passed == 1], data[data.passed == -1]


def get_judge_crosses_df(data: pd.DataFrame) -> pd.DataFrame:
    return data[["syntethic_competitor_id", "judge_name", "passed"]]
