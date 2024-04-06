# LockBox

Lockbox is a SSH file upload service that acts as a middleman, streaming the file directly without saving it on the server.

## Features

- Anonymous file uploads with no registration required.
- Zero storage on server.
- Secure and encrypted transfers.
- Fast and efficient transfers.

## Usage

1. Start the server
2. Enter command:

    ```bash
    ssh -p 2222 localhost filename < filepath
    ```

3. Share the recieved link with the user.
4. Once the user enters the link it will be streamed to them directly without saving to disk.
