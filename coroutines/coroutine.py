import asyncio

# Define a coroutine that processes a sequence of items
async def process_items(items):
    for item in items:
        print(f"Processing item: {item}")
        await asyncio.sleep(4)  # Simulate an I/O-bound operation
        print(f"Finished processing item: {item}")
        yield item  # Yield control back to the caller

# Define the main coroutine to run the process
async def main():
    items = ["apple", "banana", "cherry"]
    
    # Initialize the coroutine
    item_processor = process_items(items)
        
    
    # Re-enter the coroutine multiple times
    async for processed_item in item_processor:
        print(f"Processed item: {processed_item}")

# Run the main coroutine
asyncio.run(main())
