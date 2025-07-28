from typing import Dict, Optional
from stackelbergRefactored import StackelbergScheduler, ClusterResources

def run_stackelberg(total_cpu: float, total_memory: float, params: Optional[Dict[str, float]] = None):
    params = params or {}
    params["total_cpu"] = total_cpu
    params["total_memory"] = total_memory

    cluster = ClusterResources(params)
    scheduler = StackelbergScheduler(cluster, params)
    result = scheduler.stackelberg_equilibrium(verbose=False)

    allocations = {
        scheduler.tenants[i].name: {
            "cpu": round(cpu, 2),
            "memory": round(mem, 2),
            "replicas": replicas
        }
        for i, (cpu, mem, replicas) in enumerate(result["allocations"])
    }

    return {
        "allocations": allocations,
        "prices": {
            "cpu": round(result["prices"][0], 3),
            "memory": round(result["prices"][1], 3)
        },
        "platform_utility": result["platform_utility"],
        "metrics": {},  
        "success": True
    }
