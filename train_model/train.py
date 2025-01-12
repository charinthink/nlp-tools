from sklearn.model_selection import train_test_split
from sklearn.pipeline import Pipeline
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.linear_model import LogisticRegression
from sklearn.metrics import classification_report, accuracy_score
import pandas as pd
import joblib


class TextClassification:
    def __init__(self, df, X_test, y_test, pipeline):
        self.df = df
        self.X_test = X_test
        self.y_test = y_test
        self.pipeline = pipeline

    def report_test(self):
        y_pred = self.pipeline.predict(self.X_test)
        print("Accuracy:", accuracy_score(self.y_test, y_pred))
        print(
            "\nClassification Report:\n",
            classification_report(self.y_test, y_pred, zero_division=0),
        )


class Options:
    def __init__(self, textClassification: TextClassification):
        self.textClassification = textClassification

    def text_classfication_test(self):
        self.textClassification.report_test()


class Train:
    def __init__(self, df, X_test, y_test, pipeline):
        self.df = df
        self.X_test = X_test
        self.y_test = y_test
        self.pipeline = pipeline

    @staticmethod
    def text_classification(
        df,
        output,
        test_size=0.2,
        random_state=42,
    ):
        X = df[df.columns.values[0]]
        y = df[df.columns.values[1]]
        X_train, X_test, y_train, y_test = train_test_split(
            X, y, test_size=test_size, random_state=random_state
        )
        
        pipeline = Pipeline(
            [("vectorizer", TfidfVectorizer()), ("classifier", LogisticRegression())]
        )
        pipeline.fit(X_train, y_train)
        joblib.dump(pipeline, output)
        return Options(TextClassification(df, X_test, y_test, pipeline))
