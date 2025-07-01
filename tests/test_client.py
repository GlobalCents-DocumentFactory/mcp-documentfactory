import asyncio
import json
import os
import sys
from mcp_client import StdioServerParameters, stdio_client, ClientSession

async def main():
    if len(sys.argv) < 2:
        print("Usage: python test_client.py <file_path>")
        return

    file_path = os.path.abspath(sys.argv[1])
    if not os.path.exists(file_path):
        print(f"Error: File not found at {file_path}")
        return

    server_params = StdioServerParameters(
        command="./df-mcp-server",
        args=[],
        env={
            "NATS_URL": os.getenv("NATS_URL", "nats://127.0.0.1:4222"),
        }
    )

    async with stdio_client(server_params) as streams:
        async with ClientSession(*streams) as session:
            await session.initialize()

            print("--- Running Test Case 1: Convert to PDF ---")
            await test_convert_to_pdf(session, file_path)

            print("\n--- Running Test Case 2: Add Watermark ---")
            await test_add_watermark(session, file_path)


async def test_convert_to_pdf(session: ClientSession, file_path: str):
    pipeline_def = {
        "processors": [
            {
                "actionconvert": True,
                "settingsconvert": {"saveformat": "pdf"}
            }
        ]
    }

    result = await session.call_tool("pipeline", {
        "files": [file_path],
        "pipeline_definition": json.dumps(pipeline_def)
    })

    if result.content:
        print("Convert to PDF Result (Download URL):", result.content)
    else:
        print("Convert to PDF failed.")


async def test_add_watermark(session: ClientSession, file_path: str):
    pipeline_def = {
        "processors": [
            {
                "actionconvert": True,
                "settingsconvert": {"saveformat": "pdf"}
            },
            {
                "actionwatermark": True,
                "settingswatermarks": [
                    {
                        "watermarktype": "Text",
                        "text": "Confidential",
                        "opacity": 0.5,
                        "rotation": -45
                    }
                ]
            }
        ]
    }

    result = await session.call_tool("pipeline", {
        "files": [file_path],
        "pipeline_definition": json.dumps(pipeline_def)
    })

    if result.content:
        print("Add Watermark Result (Download URL):", result.content)
    else:
        print("Add Watermark failed.")


if __name__ == "__main__":
    asyncio.run(main())