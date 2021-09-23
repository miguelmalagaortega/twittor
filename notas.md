Podemos ver todos los remotos con

> git remote -v

Eliminar remotos con

> git remote rm nombreDelRemoto

Cargar los archivos a github por primera vez

> git init
> git add .
> git commit -m "Primer commit"
> git branch -M main
> git remote add origin git@github.com:miguelmalagaortega/twittor.git
> git push -u origin main

Cargar los archivos a heroku