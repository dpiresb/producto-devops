pipeline {
    agent { label 'aws-minikube' }

    stages {
        stage('Build') {
            steps {
                sh 'docker build -t miapp:${BRANCH_NAME}-${BUILD_NUMBER} .'
            }
        }
        stage('Load') {
            steps {
                sh 'minikube image load miapp:${BRANCH_NAME}-${BUILD_NUMBER}'
            }
        }
        stage('Desplegar en pruebas') {
            when { branch 'dev' }
            steps {
                sh '''
                    kubectl create namespace miapp-test --dry-run=client -o yaml | kubectl apply -f -
                    sed "s|image: miapp:.*|image: miapp:${BRANCH_NAME}-${BUILD_NUMBER}|" k8s/deployment.yaml | kubectl apply -n miapp-test -f -
                    kubectl apply -n miapp-test -f k8s/service.yaml
                    kubectl rollout status deployment/miapp -n miapp-test --timeout=180s
                '''
            }
        }
        stage('Desplegar en producción') {
            when { branch 'main' }
            steps {
                sh '''
                    kubectl create namespace miapp-prod --dry-run=client -o yaml | kubectl apply -f -
                    sed "s|image: miapp:.*|image: miapp:${BRANCH_NAME}-${BUILD_NUMBER}|" k8s/deployment.yaml | kubectl apply -n miapp-prod -f -
                    kubectl apply -n miapp-prod -f k8s/service.yaml
                    kubectl rollout status deployment/miapp -n miapp-prod --timeout=180s
                '''
            }
        }
    }
}