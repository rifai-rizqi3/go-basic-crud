pipeline {
    agent any
    
    stages {
        stage('Checkout') {
            steps {
                git branch: 'sonar', url: 'https://github.com/rifai-rizqi3/go-basic-crud.git'
            }
        }
        
        stage('SonarQube Analysis') {
            steps {
                withSonarQubeEnv(installationName:'sonarqube-10.5.1') {
                    sh 'sonar-scanner \
                        -Dsonar.projectKey=demo-project \
                        -Dsonar.sources=. \
                        -Dsonar.host.url=https://sonar.teamdev.id \
                        -Dsonar.token=sqp_e48848dbf36d2af9fe791d2f6658d56fdc69aea7'
                }
            }
        }
    }
}
