import pandas as pd


class LoadData:
    @staticmethod
    def csv(path, names) -> pd.DataFrame:
        df = pd.read_csv(path, sep="\t", header=None, names=names, encoding="utf-8")
        df[names[0]] = df[names[0]].str.strip()
        df[names[1]] = df[names[1].str.strip()]
        return df
