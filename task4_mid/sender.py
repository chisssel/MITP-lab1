import json
import sys

def main():
    data_to_send = {
        "user": {
            "id": 123,
            "username": "geek_user",
            "email": "geek@example.com"
        },
        "payload": "Some important data",
        "is_valid": True
    }

    json_string = json.dumps(data_to_send)

    sys.stdout.write(json_string)

if __name__ == "__main__":
    main()
