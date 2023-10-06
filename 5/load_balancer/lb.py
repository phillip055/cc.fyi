from dataclasses import dataclass
import asyncio
import requests
import itertools

@dataclass
class Server:
    host: str
    port: int
    healthy: bool = True

servers = [
    Server("localhost", 8080),
    Server("localhost", 8081),
    Server("localhost", 8082),
]

round_robin_servers = itertools.cycle(servers)

async def update_servers_health():
    async def update_server_health(server):
        while True:
            try:
                server.healthy = requests.get(server.url).status_code == 200
            except:
                server.healthy = False
            await asyncio.sleep(10)
    async with asyncio.TaskGroup() as tg:
        for server in servers:
            tg.create_task(update_server_health(server))
                

async def forward(server, data):
    reader, writer = await asyncio.open_connection(host=server.host, port=server.port)
    writer.write(data)
    await writer.drain()
    data = await reader.read(5000)
    writer.close()
    await writer.wait_closed()
    return data

async def load_balancer(reader, writer):
    server = next(round_robin_servers)
    while not server.healthy:
        server = next(server)
    data = await reader.read(5000)
    response = await forward(server, data)
    writer.write(response)
    await writer.drain()
    print("Close the connection")
    writer.close()
    await writer.wait_closed()

async def main():
    server = await asyncio.start_server(
        load_balancer, '127.0.0.1', 80)

    addrs = ', '.join(str(sock.getsockname()) for sock in server.sockets)
    print(f'Serving on {addrs}')

    async with server:
        await server.serve_forever()


if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    loop.run_until_complete(main())
    loop.run_until_complete(update_servers_health())
