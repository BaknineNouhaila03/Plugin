from fastapi import FastAPI
from pydantic import BaseModel
from typing import Dict, Optional
from stackelberg_core import run_stackelberg

app = FastAPI()

print("Creating health endpoint...")  # Debug print

@app.get("/health")
def health_check():
    print("Health endpoint called!")  # Debug print
    return {"status": "ok"}

print("Health endpoint created!")  # Debug print

class InputModel(BaseModel):
    total_cpu: float
    total_memory: float
    params: Optional[Dict[str, float]] = None

@app.post("/stackelberg/allocate")
def allocate(data: InputModel):
    return run_stackelberg(data.total_cpu, data.total_memory, data.params)

# Print all routes for debugging
@app.on_event("startup")
async def startup_event():
    print("=== All registered routes ===")
    for route in app.routes:
        print(f"Route: {route.path} - Methods: {getattr(route, 'methods', 'N/A')}")
    print("=== End routes ===")

