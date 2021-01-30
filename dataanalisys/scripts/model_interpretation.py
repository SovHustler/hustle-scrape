import typing
from more_itertools import unzip
from sklearn.linear_model import LinearRegression
import pandas as pd


class CoefData:
    coefs: list
    labels: list

    def __init__(self, coefs: list, labels: list):
        self.coefs = coefs
        self.labels = labels

    @property
    def norm_coefs(self) -> list:
        max_coef = max(self.coefs)
        return [coef / max_coef for coef in self.coefs]

    def sort(self):
        tups = sorted(list(zip(self.labels, self.coefs)), key=lambda tup: tup[1], reverse=True)

        self.labels, self.coefs = [tup[0] for tup in tups], [tup[1] for tup in tups]

        return self

    def negate(self):
        self.coefs = [-1 * coef for coef in self.coefs]
        return self


def get_diff_coef_data(coef_data: CoefData) -> CoefData:
    tups = []
    for i in range(len(coef_data.labels) // 2):
        tups.append(
            (
                coef_data.labels[i * 2],
                coef_data.coefs[i * 2] - coef_data.coefs[i * 2 + 1]
            )
        )

    tups = sorted(tups, key=lambda tup: tup[1], reverse=True)

    labels = [tup[0] for tup in tups]
    diffs = [tup[1] for tup in tups]

    return CoefData(diffs, labels)


def get_sorted_cross_based_model_coefs(reg: LinearRegression, data: pd.DataFrame) -> CoefData:
    tuples = list(zip(reg.coef_, data.columns[:-1]))

    coefs = [tup[0] for tup in tuples if tup[1] != 'total_competitors']
    labels = [tup[1] for tup in tuples if tup[1] != 'total_competitors']

    return CoefData(coefs, labels)
