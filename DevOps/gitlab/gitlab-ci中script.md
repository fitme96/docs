#  

```yaml
    - $SSH "if docker ps | grep -q ci_backend; then docker rm -f ci_backend; fi"
    - $SSH "if screen -list|grep -q $screen_name; then screen -S $screen_name -X quit; fi "
    - $SSH "docker pull hub.bugfeel.net:8443/ad-train/backend:$backend_tag"
    - $SSH "docker run -d --name ci_backend hub.bugfeel.net:8443/ad-train/backend:$backend_tag tail -F /test"
    - $SSH "ls $workdir | grep -Ev 'web_front|web_admin|venv|ad-train_uwsgi.ini|celerybeat-schedule|static|logs|local_settings.py' |awk -v workdir="$workdir" '{print workdir \"/\"  \$0}'|xargs -I {} rm -rf {}"
    - $SSH "docker cp ci_backend:/backend/. $workdir/"
    - $SSH "cp $workdir/local_settings.py $workdir/ad_train/"
    - $SSH "docker rm -f ci_backend"
    - $SSH "cd $workdir && source venv/bin/activate && python manage.py migrate && python manage.py collectstatic --noinput"
    - $SSH "screen -S $screen_name -X quit"
    - $SSH "cd $workdir && source venv/bin/activate && screen -dmS $screen_name python manage.py runserver 0.0.0.0:$port"



```
