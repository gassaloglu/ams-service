pipeline {
    agent {
        kubernetes {
            inheritFrom 'kaniko'
        }
    }

    options {
        buildDiscarder logRotator(artifactDaysToKeepStr: '', artifactNumToKeepStr: '5', daysToKeepStr: '', numToKeepStr: '5')
    }

    tools{
        go '1.23.2'
            //dockerTool 'latest'
    }
    stages {

        stage('Git'){
            steps{
                // Add github project
                git branch: 'main', credentialsId: 'test-username', url: 'https://github.com/airline-management-system/ams-service.git'
            }
        }

        // stage('Unit Tests'){
        //     steps {
        //         // Run Go unit tests
        //         sh 'go test -v ./...'
        //     }
        // }


        stage('Docker Build & Push'){
            environment {
                PATH = "/busybox:/kaniko:$PATH"
            }
            steps {
                container(name: 'kaniko', shell: '/busybox/sh') {
                    sh '/kaniko/executor --dockerfile `pwd`/Dockerfile --context `pwd` --destination=harbor.turkey-diminished.ts.net/ams/ams-service:latest'
                }
            }
        }
    }
}
