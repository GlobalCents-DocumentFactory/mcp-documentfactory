# Testing the DocumentFactory MCP Server

This directory contains a Python script to test the DocumentFactory MCP server.

## Prerequisites

- Python 3.6+
- A running NATS server.
- The core DocumentFactory service must be running.

## Setup

1.  **Install Dependencies:**

    Install the required Python packages using `pip` and the `requirements.txt` file.

    ```bash
    pip install -r requirements.txt
    ```

2.  **Build the MCP Server:**

    Before running the tests, make sure you have the latest version of the MCP server built.

    ```bash
    go build -o df-mcp-server ./cmd/df-mcp-server
    ```

3.  **Prepare a Test File:**

    Create a sample file to be used as input for the tests.

    ```bash
    echo "This is a test." > test.txt
    ```

## Running the Tests

Execute the `test_client.py` script, passing the path to your test file as an argument.

```bash
python test_client.py /path/to/your/test.txt
```

*Replace `/path/to/your/test.txt` with the absolute path to the test file you created.*

The script will execute two test cases:
1.  Convert the input file to a PDF.
2.  Convert the input file to a PDF and add a "Confidential" watermark.

## Verification

The test script will print download URLs for the output of each test case. You can use `curl` or your web browser to download the resulting files and verify their contents.

**Example:**

```bash
# For the PDF conversion test
curl -o output.pdf "http://localhost:8080/download/some-token-1"

# For the watermarking test
curl -o watermarked_output.pdf "http://localhost:8080/download/some-token-2"
```
