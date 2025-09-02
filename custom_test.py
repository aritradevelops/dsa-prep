import time


class TestCase:
    def __init__(self, inputs, expected):
        self.inputs = inputs
        self.expected = expected


class TestRunner:
    def __init__(self, func, test_cases):
        self.func = func
        self.test_cases = test_cases

    def run(self):
        print(
            f"Testing {self.func.__name__} with {len(self.test_cases)} test cases.\n{'-'*80}")
        total_time = 0.0
        for i, case in enumerate(self.test_cases, 1):
            start_time = time.time()
            try:
                result = self.func(*case.inputs)
                success = result == case.expected
            except Exception as e:
                result = f"Exception: {e}"
                success = False
            end_time = time.time()
            elapsed = end_time - start_time
            total_time += elapsed
            status = "✅ PASS" if success else "❌ FAIL"
            print(
                f"Test {i}: Inputs={case.inputs} | Expected={case.expected} | Got={result} | {status} | Time={elapsed:.6f} seconds")
        print('-' * 80)
        print(f"Total Runtime for all tests: {total_time:.6f} seconds")
        print('-' * 80)
