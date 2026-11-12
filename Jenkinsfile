pipeline {
    agent { label 'aws-minikube' }

    stages {
        stage('Build') {
            steps {
                sh 'docker build -t miapp:$BUILD_NUMBER .'
            }
        }
        stage('Load') {
            steps {
                sh 'minikube image load miapp:$BUILD_NUMBER'
            }
        }
        stage('Deploy') {
            steps {
                sh '''
                    sed "s|image: miapp:.*|image: miapp:$BUILD_NUMBER|" k8s/deployment.yaml | kubectl apply -f -
                    kubectl apply -f k8s/service.yaml
                    kubectl rollout status deployment/miapp --timeout=180s
                '''
            }
        }
    }
}