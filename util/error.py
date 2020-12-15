class ConnectError(Exception):
    def __init__(self, e):
        self.exception = e

    def __str__(self):
        return str(self.exception)
