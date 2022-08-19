import aioboto3

async def service_list(region_name:str, service_name:str=None) -> list:
    """Uses the region passed to return all available services."""
    session = aioboto3.Session(region_name=region_name)
    services:list = session.get_available_services()

    return [service_name] if service_name in services else services
