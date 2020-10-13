node {
    stage('Git Pull') {
        git 'https://github.com/ShroXd/go-server.git'
    }

    stage('Build Image') {
        try {
            sh 'docker stop go-server'
            sh 'docker rm go-server'
            sh 'docker rmi go-server'
        }
        catch (exc) {
            echo '无需清理容器残余'
        }
        finally {
            sh 'docker build -t go-server:latest .'
        }
    }

    stage('Deploy') {
        try {
          sh 'docker run -p 1451:8080 -v /home/Shroud/conf/go-server.ini:/code/conf/app.ini -v /home/Shroud/logs/:/code/tmp/logs/ --name go-server -d go-server'
        }
        catch (exc) {
          echo '运行容器失败'
          sh 'docker stop go-server'
          sh 'docker rm go-server'
          sh 'docker run -p 1451:8080 -v /home/Shroud/conf/go-server.ini:/code/conf/app.ini -v /home/Shroud/logs/:/code/tmp/logs/ --name go-server -d go-server'
        }
    }

}
