import unittest
import json
import sys
from io import StringIO
from sender import main

class TestSender(unittest.TestCase):

    def test_main_output(self):
        original_stdout = sys.stdout
        captured_output = StringIO()
        sys.stdout = captured_output

        main()

        sys.stdout = original_stdout

        output_str = captured_output.getvalue()

        self.assertTrue(len(output_str) > 0, "Вывод не должен быть пустым")

        try:
            data = json.loads(output_str)
        except json.JSONDecodeError:
            self.fail("Вывод не является валидным JSON")

        self.assertIn("user", data)
        self.assertIn("payload", data)
        self.assertEqual(data["user"]["id"], 123)
        self.assertEqual(data["user"]["username"], "geek_user")
        self.assertEqual(data["is_valid"], True)

if __name__ == '__main__':
    unittest.main()

