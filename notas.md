Podemos ver todos los remotos con

> git remote -v

Eliminar remotos con

> git remote rm nombreDelRemoto

***NOTA:*** en la aplicacion hay que tener cuidado en que rama guardamos todo

1. El git se guardo en la rama "branch" main, asi que todo debe ir en esa rama
2. Tener cuidado ya que en algunas aplicaciones como heroku aparece la rama master

Cargar los archivos a github por primera vez

> git init
> git add .
> git commit -m "Primer commit"
> git branch -M main
> git remote add origin git@github.com:miguelmalagaortega/twittor.git
> git push -u origin main

Cargar los archivos a heroku por primera vez

> heroku login
> heroku git:remote -a twittordcn
> git push heroku main

4. Ahora al hacer las cargas hacia los repositorios haremos los siguiente

4.1. Para Github
> git push -u origin main
4.2. Para heroku
> git push heroku main

