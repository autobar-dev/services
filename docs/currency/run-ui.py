import os
import subprocess
import signal

with open("pid", "r") as f:
    data = f.read().strip()
    
    if data:
        os.kill(int(data), signal.SIGKILL)

proc = subprocess.Popen(['swagger', 'serve', '-F=redoc', '--no-open', '-p=9092', 'swagger.json'])

with open("pid", "w") as f:
    f.write(str(proc.pid))
