import os

class A:
    pass

class SomeClass(A):
    def __init__(self) -> None:
        super().__init__()

    def some_method(self):
        print("some_method")

    def some_method2(self, a, b):
        """ some_method2 """
        return a + b


if __name__ == '__main__':
    var = SomeClass()
    var.some_method()
    print(var.some_method2(1, 2))
