import typing
import pandas as pd
from sklearn.linear_model import LassoCV
from sklearn.model_selection import train_test_split


def get_regression(data: pd.DataFrame) -> typing.Tuple[LassoCV, float]:
    x_train, x_test, y_train, y_test = train_test_split(data[data.columns[:-1]], data[data.columns[-1]], test_size=0.1, random_state=42)

    reg = LassoCV(cv=5, random_state=42).fit(x_train, y_train)

    return reg, reg.score(x_test, y_test)
