import asyncio

async def slow_coroutine(name, delay):
    print(f"{name} started")
    await asyncio.sleep(delay)  # Simulate a blocking I/O operation
    print(f"{name} completed after {delay} seconds")

async def fast_coroutine(name):
    print(f"{name} started")
    for i in range(10):
        print(f"{name} iteration {i}")
        await asyncio.sleep(1)  # Yield control back to the event loop
    print(f"{name} completed")

async def main():
    # Run both coroutines concurrently
    
    ### There is No True Preemption/Control-of-Interuption: While fast_coroutine is frequently yielding control, once slow_coroutine is ready to 
    #######  finish (after its sleep), it will execute the rest of its code until it hits the next await or until it finishes.
    await asyncio.gather(
        slow_coroutine("SlowTask", 5),   #
        fast_coroutine("FastTask")
    )

# Run the main coroutine
asyncio.run(main())
