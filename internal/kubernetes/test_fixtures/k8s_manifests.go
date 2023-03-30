package kubernetes

var EXAMPLE_DEPLOYMENT = `
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    argocd.argoproj.io/sync-wave: '3'
    cnrm.cloud.google.com/project-id: test-prod-customers
    deployment.kubernetes.io/revision: '129'
    kubectl.kubernetes.io/last-applied-configuration: >
      {"apiVersion":"apps/v1","kind":"Deployment","metadata":{"annotations":{"argocd.argoproj.io/sync-wave":"3","cnrm.cloud.google.com/project-id":"test-prod-customers"},"labels":{"app.kubernetes.io/component":"test-product-deployment","app.kubernetes.io/instance":"prod-customer-data","app.kubernetes.io/name":"customer-data","app.kubernetes.io/part-of":"events","k8s.test.com/team":"data-platform","tags.datadoghq.com/env":"prod","tags.datadoghq.com/service":"customer-data","tags.datadoghq.com/version":"56869a5"},"name":"customer-data","namespace":"customer-data"},"spec":{"selector":{"matchLabels":{"app":"customer-data","app.kubernetes.io/part-of":"events","k8s.test.com/team":"data-platform","tags.datadoghq.com/env":"prod"}},"strategy":{"rollingUpdate":{"maxSurge":"10%","maxUnavailable":"0%"},"type":"RollingUpdate"},"template":{"metadata":{"annotations":{"cnrm.cloud.google.com/project-id":"test-prod-customers","prometheus.io/port":"9102","prometheus.io/scrape":"true","sidecar.istio.io/capNetBindService":"true","sidecar.istio.io/userVolume":"[{\"name\":\"descriptor\",
      \"secret\":{\"secretName\":\"descriptors-pb\"}}]","sidecar.istio.io/userVolumeMount":"[{\"name\":\"descriptor\",
      \"mountPath\":\"/etc/envoy\",
      \"readonly\":true}]"},"labels":{"app":"customer-data","app.kubernetes.io/component":"test-product-deployment","app.kubernetes.io/name":"customer-data","app.kubernetes.io/part-of":"events","k8s.test.com/team":"data-platform","tags.datadoghq.com/env":"prod","tags.datadoghq.com/service":"customer-data","tags.datadoghq.com/version":"56869a5"}},"spec":{"containers":[{"args":["server"],"env":[{"name":"SERVICE_NAME","value":"customer-data"},{"name":"HOST_IP","valueFrom":{"fieldRef":{"fieldPath":"status.hostIP"}}},{"name":"TRACE_ADDRESS","value":"$(HOST_IP):4317"},{"name":"BLACKLIST_FILE","value":"/app/blacklist.yaml"},{"name":"DD_ENV","valueFrom":{"fieldRef":{"fieldPath":"metadata.labels['tags.datadoghq.com/env']"}}},{"name":"DD_SERVICE","valueFrom":{"fieldRef":{"fieldPath":"metadata.labels['tags.datadoghq.com/service']"}}},{"name":"DD_VERSION","valueFrom":{"fieldRef":{"fieldPath":"metadata.labels['tags.datadoghq.com/version']"}}},{"name":"GARBAGE","value":"delete-me-later"}],"envFrom":[{"configMapRef":{"name":"config-86t98h48c7"}},{"configMapRef":{"name":"rpc-config-ht9m76d66k"}},{"secretRef":{"name":"customer-data-secrets"}}],"image":"us.gcr.io/test-ops/apps/customer-data:56869a5","imagePullPolicy":"Always","livenessProbe":{"exec":{"command":["/bin/grpc_health_probe","-addr=:50051"]},"initialDelaySeconds":10,"periodSeconds":10},"name":"customer-data","ports":[{"containerPort":50051,"name":"grpc"},{"containerPort":8080,"name":"http"},{"containerPort":9102,"name":"http-metrics"}],"readinessProbe":{"exec":{"command":["/bin/grpc_health_probe","-addr=:50051"]},"initialDelaySeconds":5},"resources":{"limits":{"memory":"2Gi"},"requests":{"cpu":"2","memory":"1Gi"}},"volumeMounts":[{"mountPath":"/app","name":"blacklist"}]}],"dnsConfig":{"options":[{"name":"ndots","value":"3"}]},"serviceAccountName":"argo-sa-customer-data","volumes":[{"configMap":{"name":"blacklist"},"name":"blacklist"}]}}}}
  creationTimestamp: '2021-08-24T19:06:44Z'
  generation: 7204
  labels:
    app.kubernetes.io/component: test-product-deployment
    app.kubernetes.io/instance: prod-customer-data
    app.kubernetes.io/name: customer-data
    app.kubernetes.io/part-of: events
    k8s.test.com/team: data-platform
    tags.datadoghq.com/env: prod
    tags.datadoghq.com/service: customer-data
    tags.datadoghq.com/version: 56869a5
  name: customer-data
  namespace: customer-data
  resourceVersion: '4020930045'
  uid: 09026a6a-3310-4623-8ce2-f1dd310c9391
spec:
  progressDeadlineSeconds: 600
  replicas: 10
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      app: customer-data
      app.kubernetes.io/part-of: events
      k8s.test.com/team: data-platform
      tags.datadoghq.com/env: prod
  strategy:
    rollingUpdate:
      maxSurge: 10%
      maxUnavailable: 0%
    type: RollingUpdate
  template:
    metadata:
      annotations:
        cnrm.cloud.google.com/project-id: test-prod-customers
        kubectl.kubernetes.io/restartedAt: '2022-09-06T20:46:32-04:00'
        prometheus.io/port: '9102'
        prometheus.io/scrape: 'true'
        sidecar.istio.io/capNetBindService: 'true'
        sidecar.istio.io/userVolume: '[{"name":"descriptor", "secret":{"secretName":"descriptors-pb"}}]'
        sidecar.istio.io/userVolumeMount: '[{"name":"descriptor", "mountPath":"/etc/envoy", "readonly":true}]'
      creationTimestamp: null
      labels:
        app: customer-data
        app.kubernetes.io/component: test-product-deployment
        app.kubernetes.io/name: customer-data
        app.kubernetes.io/part-of: events
        k8s.test.com/team: data-platform
        tags.datadoghq.com/env: prod
        tags.datadoghq.com/service: customer-data
        tags.datadoghq.com/version: 56869a5
    spec:
      containers:
        - args:
            - server
          env:
            - name: SERVICE_NAME
              value: customer-data
            - name: HOST_IP
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: status.hostIP
            - name: TRACE_ADDRESS
              value: '$(HOST_IP):4317'
            - name: BLACKLIST_FILE
              value: /app/blacklist.yaml
            - name: DD_ENV
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: 'metadata.labels[''tags.datadoghq.com/env'']'
            - name: DD_SERVICE
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: 'metadata.labels[''tags.datadoghq.com/service'']'
            - name: DD_VERSION
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: 'metadata.labels[''tags.datadoghq.com/version'']'
            - name: GARBAGE
              value: delete-me-later
          envFrom:
            - configMapRef:
                name: config-86t98h48c7
            - configMapRef:
                name: rpc-config-ht9m76d66k
            - secretRef:
                name: customer-data-secrets
          image: 'us.gcr.io/test-ops/apps/customer-data:56869a5'
          imagePullPolicy: Always
          livenessProbe:
            exec:
              command:
                - /bin/grpc_health_probe
                - '-addr=:50051'
            failureThreshold: 3
            initialDelaySeconds: 10
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 1
          name: customer-data
          ports:
            - containerPort: 50051
              name: grpc
              protocol: TCP
            - containerPort: 8080
              name: http
              protocol: TCP
            - containerPort: 9102
              name: http-metrics
              protocol: TCP
          readinessProbe:
            exec:
              command:
                - /bin/grpc_health_probe
                - '-addr=:50051'
            failureThreshold: 3
            initialDelaySeconds: 5
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 1
          resources:
            limits:
              memory: 2Gi
            requests:
              cpu: '2'
              memory: 1Gi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
            - mountPath: /app
              name: blacklist
      dnsConfig:
        options:
          - name: ndots
            value: '3'
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      serviceAccount: argo-sa-customer-data
      serviceAccountName: argo-sa-customer-data
      terminationGracePeriodSeconds: 30
      volumes:
        - configMap:
            defaultMode: 420
            name: blacklist
          name: blacklist
status:
  availableReplicas: 10
  conditions:
    - lastTransitionTime: '2021-08-24T19:40:12Z'
      lastUpdateTime: '2022-11-11T21:40:09Z'
      message: ReplicaSet "customer-data-544fb5d4c6" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: 'True'
      type: Progressing
    - lastTransitionTime: '2022-11-15T15:09:29Z'
      lastUpdateTime: '2022-11-15T15:09:29Z'
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: 'True'
      type: Available
  observedGeneration: 7204
  readyReplicas: 10
  replicas: 10
  updatedReplicas: 10
`

var PARSED_JSON = `[{"apiVersion":"apps/v1","kind":"Deployment","metadata":{"annotations":{"argocd.argoproj.io/sync-wave":"3","cnrm.cloud.google.com/project-id":"test-prod-customers","deployment.kubernetes.io/revision":"129","kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"apps/v1\",\"kind\":\"Deployment\",\"metadata\":{\"annotations\":{\"argocd.argoproj.io/sync-wave\":\"3\",\"cnrm.cloud.google.com/project-id\":\"test-prod-customers\"},\"labels\":{\"app.kubernetes.io/component\":\"test-product-deployment\",\"app.kubernetes.io/instance\":\"prod-customer-data\",\"app.kubernetes.io/name\":\"customer-data\",\"app.kubernetes.io/part-of\":\"events\",\"k8s.test.com/team\":\"data-platform\",\"tags.datadoghq.com/env\":\"prod\",\"tags.datadoghq.com/service\":\"customer-data\",\"tags.datadoghq.com/version\":\"56869a5\"},\"name\":\"customer-data\",\"namespace\":\"customer-data\"},\"spec\":{\"selector\":{\"matchLabels\":{\"app\":\"customer-data\",\"app.kubernetes.io/part-of\":\"events\",\"k8s.test.com/team\":\"data-platform\",\"tags.datadoghq.com/env\":\"prod\"}},\"strategy\":{\"rollingUpdate\":{\"maxSurge\":\"10%\",\"maxUnavailable\":\"0%\"},\"type\":\"RollingUpdate\"},\"template\":{\"metadata\":{\"annotations\":{\"cnrm.cloud.google.com/project-id\":\"test-prod-customers\",\"prometheus.io/port\":\"9102\",\"prometheus.io/scrape\":\"true\",\"sidecar.istio.io/capNetBindService\":\"true\",\"sidecar.istio.io/userVolume\":\"[{\name\:\descriptor\, \secret\:{\secretName\:\descriptors-pb\}}]\",\"sidecar.istio.io/userVolumeMount\":\"[{\name\:\descriptor\, \mountPath\:\/etc/envoy\, \readonly\:true}]\"},\"labels\":{\"app\":\"customer-data\",\"app.kubernetes.io/component\":\"test-product-deployment\",\"app.kubernetes.io/name\":\"customer-data\",\"app.kubernetes.io/part-of\":\"events\",\"k8s.test.com/team\":\"data-platform\",\"tags.datadoghq.com/env\":\"prod\",\"tags.datadoghq.com/service\":\"customer-data\",\"tags.datadoghq.com/version\":\"56869a5\"}},\"spec\":{\"containers\":[{\"args\":[\"server\"],\"env\":[{\"name\":\"SERVICE_NAME\",\"value\":\"customer-data\"},{\"name\":\"HOST_IP\",\"valueFrom\":{\"fieldRef\":{\"fieldPath\":\"status.hostIP\"}}},{\"name\":\"TRACE_ADDRESS\",\"value\":\"$(HOST_IP):4317\"},{\"name\":\"BLACKLIST_FILE\",\"value\":\"/app/blacklist.yaml\"},{\"name\":\"DD_ENV\",\"valueFrom\":{\"fieldRef\":{\"fieldPath\":\"metadata.labels['tags.datadoghq.com/env']\"}}},{\"name\":\"DD_SERVICE\",\"valueFrom\":{\"fieldRef\":{\"fieldPath\":\"metadata.labels['tags.datadoghq.com/service']\"}}},{\"name\":\"DD_VERSION\",\"valueFrom\":{\"fieldRef\":{\"fieldPath\":\"metadata.labels['tags.datadoghq.com/version']\"}}},{\"name\":\"GARBAGE\",\"value\":\"delete-me-later\"}],\"envFrom\":[{\"configMapRef\":{\"name\":\"config-86t98h48c7\"}},{\"configMapRef\":{\"name\":\"rpc-config-ht9m76d66k\"}},{\"secretRef\":{\"name\":\"customer-data-secrets\"}}],\"image\":\"us.gcr.io/test-ops/apps/customer-data:56869a5\",\"imagePullPolicy\":\"Always\",\"livenessProbe\":{\"exec\":{\"command\":[\"/bin/grpc_health_probe\",\"-addr=:50051\"]},\"initialDelaySeconds\":10,\"periodSeconds\":10},\"name\":\"customer-data\",\"ports\":[{\"containerPort\":50051,\"name\":\"grpc\"},{\"containerPort\":8080,\"name\":\"http\"},{\"containerPort\":9102,\"name\":\"http-metrics\"}],\"readinessProbe\":{\"exec\":{\"command\":[\"/bin/grpc_health_probe\",\"-addr=:50051\"]},\"initialDelaySeconds\":5},\"resources\":{\"limits\":{\"memory\":\"2Gi\"},\"requests\":{\"cpu\":\"2\",\"memory\":\"1Gi\"}},\"volumeMounts\":[{\"mountPath\":\"/app\",\"name\":\"blacklist\"}]}],\"dnsConfig\":{\"options\":[{\"name\":\"ndots\",\"value\":\"3\"}]},\"serviceAccountName\":\"argo-sa-customer-data\",\"volumes\":[{\"configMap\":{\"name\":\"blacklist\"},\"name\":\"blacklist\"}]}}}}\n"},"creationTimestamp":"2021-08-24T19:06:44Z","generation":7204,"labels":{"app.kubernetes.io/component":"test-product-deployment","app.kubernetes.io/instance":"prod-customer-data","app.kubernetes.io/name":"customer-data","app.kubernetes.io/part-of":"events","k8s.test.com/team":"data-platform","tags.datadoghq.com/env":"prod","tags.datadoghq.com/service":"customer-data","tags.datadoghq.com/version":"56869a5"},"name":"customer-data","namespace":"customer-data","resourceVersion":"4020930045","uid":"09026a6a-3310-4623-8ce2-f1dd310c9391"},"spec":{"progressDeadlineSeconds":600,"replicas":10,"revisionHistoryLimit":10,"selector":{"matchLabels":{"app":"customer-data","app.kubernetes.io/part-of":"events","k8s.test.com/team":"data-platform","tags.datadoghq.com/env":"prod"}},"strategy":{"rollingUpdate":{"maxSurge":"10%","maxUnavailable":"0%"},"type":"RollingUpdate"},"template":{"metadata":{"annotations":{"cnrm.cloud.google.com/project-id":"test-prod-customers","kubectl.kubernetes.io/restartedAt":"2022-09-06T20:46:32-04:00","prometheus.io/port":"9102","prometheus.io/scrape":"true","sidecar.istio.io/capNetBindService":"true","sidecar.istio.io/userVolume":"[{\"name\":\"descriptor\", \"secret\":{\"secretName\":\"descriptors-pb\"}}]","sidecar.istio.io/userVolumeMount":"[{\"name\":\"descriptor\", \"mountPath\":\"/etc/envoy\", \"readonly\":true}]"},"creationTimestamp":null,"labels":{"app":"customer-data","app.kubernetes.io/component":"test-product-deployment","app.kubernetes.io/name":"customer-data","app.kubernetes.io/part-of":"events","k8s.test.com/team":"data-platform","tags.datadoghq.com/env":"prod","tags.datadoghq.com/service":"customer-data","tags.datadoghq.com/version":"56869a5"}},"spec":{"containers":[{"args":["server"],"env":[{"name":"SERVICE_NAME","value":"customer-data"},{"name":"HOST_IP","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"status.hostIP"}}},{"name":"TRACE_ADDRESS","value":"$(HOST_IP):4317"},{"name":"BLACKLIST_FILE","value":"/app/blacklist.yaml"},{"name":"DD_ENV","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"metadata.labels['tags.datadoghq.com/env']"}}},{"name":"DD_SERVICE","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"metadata.labels['tags.datadoghq.com/service']"}}},{"name":"DD_VERSION","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"metadata.labels['tags.datadoghq.com/version']"}}},{"name":"GARBAGE","value":"delete-me-later"}],"envFrom":[{"configMapRef":{"name":"config-86t98h48c7"}},{"configMapRef":{"name":"rpc-config-ht9m76d66k"}},{"secretRef":{"name":"customer-data-secrets"}}],"image":"us.gcr.io/test-ops/apps/customer-data:56869a5","imagePullPolicy":"Always","livenessProbe":{"exec":{"command":["/bin/grpc_health_probe","-addr=:50051"]},"failureThreshold":3,"initialDelaySeconds":10,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":1},"name":"customer-data","ports":[{"containerPort":50051,"name":"grpc","protocol":"TCP"},{"containerPort":8080,"name":"http","protocol":"TCP"},{"containerPort":9102,"name":"http-metrics","protocol":"TCP"}],"readinessProbe":{"exec":{"command":["/bin/grpc_health_probe","-addr=:50051"]},"failureThreshold":3,"initialDelaySeconds":5,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":1},"resources":{"limits":{"memory":"2Gi"},"requests":{"cpu":"2","memory":"1Gi"}},"terminationMessagePath":"/dev/termination-log","terminationMessagePolicy":"File","volumeMounts":[{"mountPath":"/app","name":"blacklist"}]}],"dnsConfig":{"options":[{"name":"ndots","value":"3"}]},"dnsPolicy":"ClusterFirst","restartPolicy":"Always","schedulerName":"default-scheduler","securityContext":{},"serviceAccount":"argo-sa-customer-data","serviceAccountName":"argo-sa-customer-data","terminationGracePeriodSeconds":30,"volumes":[{"configMap":{"defaultMode":420,"name":"blacklist"},"name":"blacklist"}]}}},"status":{"availableReplicas":10,"conditions":[{"lastTransitionTime":"2021-08-24T19:40:12Z","lastUpdateTime":"2022-11-11T21:40:09Z","message":"ReplicaSet \"customer-data-544fb5d4c6\" has successfully progressed.","reason":"NewReplicaSetAvailable","status":"True","type":"Progressing"},{"lastTransitionTime":"2022-11-15T15:09:29Z","lastUpdateTime":"2022-11-15T15:09:29Z","message":"Deployment has minimum availability.","reason":"MinimumReplicasAvailable","status":"True","type":"Available"}],"observedGeneration":7204,"readyReplicas":10,"replicas":10,"updatedReplicas":10}}]`

var EXAMPLE_LIST = `
apiVersion: v1
items:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "5"
    creationTimestamp: "2022-07-18T19:54:14Z"
    generation: 5
    labels:
      run: airflow-scheduler
    name: airflow-scheduler
    namespace: composer-2-0-19-airflow-2-2-5-e4894b1f
    resourceVersion: "304279318"
    uid: 19fa28fe-2a90-40b2-9f2f-ed7b7a369a2d
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        run: airflow-scheduler
    strategy:
      rollingUpdate:
        maxSurge: 1
        maxUnavailable: 1
      type: RollingUpdate
    template:
      metadata:
        annotations:
          cluster-autoscaler.kubernetes.io/safe-to-evict: "true"
        creationTimestamp: null
        labels:
          composer-system-pod: "true"
          config_id: 34c6b681-2b2a-4726-b5b3-59b20bc435a6
          run: airflow-scheduler
      spec:
        containers:
        - args:
          - scheduler
          env:
          - name: CLOUDSDK_METRICS_ENVIRONMENT
            value: 2.2.5+composer
          - name: GCS_BUCKET
            value: us-central1-airflow-bed-152ea37f-bucket
          - name: AIRFLOW_HOME
            value: /etc/airflow
          - name: DAGS_FOLDER
            value: /home/airflow/gcs/dags
          - name: SQL_HOST
            value: airflow-sqlproxy-service.composer-system.svc.cluster.local
          - name: SQL_DATABASE
            value: composer-2-0-19-airflow-2-2-5-e4894b1f
          - name: SQL_USER
            value: root
          - name: SQL_PASSWORD
            valueFrom:
              secretKeyRef:
                key: sql_password
                name: airflow-secrets
          - name: GCSFUSE_EXTRACTED
            value: "TRUE"
          - name: COMPOSER_VERSION
            value: 2.0.19
          - name: AIRFLOW__WEBSERVER__BASE_URL
            value: https://2bddbf62981044cc9b7df0b2fb317261-dot-us-central1.composer.googleusercontent.com
          - name: AIRFLOW__CORE__SQL_ALCHEMY_CONN
            value: postgresql+psycopg2://$(SQL_USER):$(SQL_PASSWORD)@airflow-sqlproxy-service.composer-system.svc.cluster.local:3306/$(SQL_DATABASE)
          - name: AIRFLOW__CORE__FERNET_KEY
            valueFrom:
              secretKeyRef:
                key: fernet_key
                name: airflow-secrets
          - name: GCP_PROJECT
            value: digraph-2021
          - name: COMPOSER_LOCATION
            value: us-central1
          - name: COMPOSER_GKE_ZONE
          - name: COMPOSER_GKE_NAME
            value: us-central1-airflow-bed-152ea37f-gke
          - name: AUTOGKE
            value: "TRUE"
          - name: COMPOSER_GKE_LOCATION
            value: us-central1
          - name: COMPOSER_PYTHON_VERSION
            value: "3"
          - name: COMPOSER_ENVIRONMENT
            value: airflow-bed
          - name: COMPOSER_VERSIONED_NAMESPACE
            value: composer-2-0-19-airflow-2-2-5-e4894b1f
          - name: GKE_CLUSTER_NAME
            value: us-central1-airflow-bed-152ea37f-gke
          - name: POD_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.name
          - name: CONTAINER_NAME
            value: airflow-scheduler
          image: us-central1-docker.pkg.dev/digraph-2021/composer-images-us-central1-airflow-bed-152ea37f-gke/33e72803-04e7-4edc-ab9c-b7e19661b109
          imagePullPolicy: IfNotPresent
          livenessProbe:
            exec:
              command:
              - /var/local/scheduler_checker.py
            failureThreshold: 3
            initialDelaySeconds: 120
            periodSeconds: 60
            successThreshold: 1
            timeoutSeconds: 30
          name: airflow-scheduler
          resources:
            limits:
              cpu: 850m
              ephemeral-storage: 921Mi
              memory: 1632Mi
            requests:
              cpu: 850m
              ephemeral-storage: 921Mi
              memory: 1632Mi
          securityContext:
            capabilities:
              drop:
              - NET_RAW
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /etc/airflow/airflow_cfg
            name: airflow-config
          - mountPath: /home/airflow/gcs
            name: gcsdir
          - mountPath: /home/airflow/container-comms
            name: container-comms
          - mountPath: /home/airflow/gcsfuse
            mountPropagation: HostToContainer
            name: gcsfuse
        - args:
          - /home/airflow/gcs
          env:
          - name: GCS_BUCKET
            value: us-central1-airflow-bed-152ea37f-bucket
          - name: SQL_DATABASE
            value: composer-2-0-19-airflow-2-2-5-e4894b1f
          - name: SQL_USER
            value: root
          - name: SQL_PASSWORD
            valueFrom:
              secretKeyRef:
                key: sql_password
                name: airflow-secrets
          - name: COMPOSER_GKE_ZONE
          - name: COMPOSER_GKE_NAME
            value: us-central1-airflow-bed-152ea37f-gke
          - name: AUTOGKE
            value: "TRUE"
          - name: COMPOSER_GKE_LOCATION
            value: us-central1
          image: us-docker.pkg.dev/cloud-airflow-releaser/gcs-syncd/gcs-syncd:cloud_composer_service_2022-06-28-RC0
          imagePullPolicy: IfNotPresent
          name: gcs-syncd
          resources:
            limits:
              cpu: 150m
              ephemeral-storage: 102Mi
              memory: 288Mi
            requests:
              cpu: 150m
              ephemeral-storage: 102Mi
              memory: 288Mi
          securityContext:
            capabilities:
              drop:
              - NET_RAW
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /home/airflow/gcs
            name: gcsdir
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          seccompProfile:
            type: RuntimeDefault
        terminationGracePeriodSeconds: 30
        volumes:
        - configMap:
            defaultMode: 420
            name: airflow-configmap
          name: airflow-config
        - emptyDir: {}
          name: gcsdir
        - hostPath:
            path: /var/composer/gcs_mount_status
            type: ""
          name: container-comms
        - hostPath:
            path: /var/composer/gcs_mount
            type: ""
          name: gcsfuse
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-07-18T19:54:14Z"
      lastUpdateTime: "2022-07-18T19:54:14Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    - lastTransitionTime: "2022-07-18T19:54:14Z"
      lastUpdateTime: "2022-09-07T22:51:51Z"
      message: ReplicaSet "airflow-scheduler-84d7877c6d" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    observedGeneration: 5
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "4"
    creationTimestamp: "2022-07-18T19:54:14Z"
    generation: 4
    labels:
      run: airflow-webserver
    name: airflow-webserver
    namespace: composer-2-0-19-airflow-2-2-5-e4894b1f
    resourceVersion: "284918896"
    uid: c676f566-df02-4d51-8ab1-e2ed39ace0e3
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        run: airflow-webserver
    strategy:
      rollingUpdate:
        maxSurge: 1
        maxUnavailable: 0
      type: RollingUpdate
    template:
      metadata:
        creationTimestamp: null
        labels:
          config_id: 9223150f-a6e3-4a55-a110-9a3f8490f14e
          run: airflow-webserver
      spec:
        containers:
        - args:
          - webserver
          env:
          - name: AIRFLOW_WEBSERVER
            value: "True"
          - name: SQL_HOST
            value: airflow-sqlproxy-service.composer-system.svc.cluster.local
          - name: SQL_DATABASE
            value: composer-2-0-19-airflow-2-2-5-e4894b1f
          - name: SQL_USER
            value: root
          - name: SQL_PASSWORD
            valueFrom:
              secretKeyRef:
                key: sql_password
                name: airflow-secrets
          - name: COMPOSER_VERSION
            value: 2.0.19
          - name: AIRFLOW__WEBSERVER__BASE_URL
            value: https://2bddbf62981044cc9b7df0b2fb317261-dot-us-central1.composer.googleusercontent.com
          - name: AIRFLOW__WEBSERVER__JWT_PUBLIC_KEY_URL
            value: https://us-central1.composer.cloud.google.com/jwt-public-key
          - name: AIRFLOW__WEBSERVER__INVERTING_PROXY_BACKEND_ID
            value: 2bddbf62981044cc9b7df0b2fb317261
          - name: AIRFLOW__WEBSERVER__SECRET_KEY
            value: some-random-id
          - name: AIRFLOW__WEBSERVER__UPDATE_FAB_PERMS
            value: "True"
          - name: AIRFLOW__CORE__FERNET_KEY
            valueFrom:
              secretKeyRef:
                key: fernet_key
                name: airflow-secrets
          - name: AIRFLOW__CORE__SQL_ALCHEMY_CONN
            value: postgresql+psycopg2://$(SQL_USER):$(SQL_PASSWORD)@airflow-sqlproxy-service.composer-system.svc.cluster.local:3306/$(SQL_DATABASE)
          - name: CLOUDSDK_METRICS_ENVIRONMENT
            value: 2.2.5+composer
          - name: GCS_BUCKET
            value: us-central1-airflow-bed-152ea37f-bucket
          - name: AIRFLOW_HOME
            value: /etc/airflow
          - name: DAGS_FOLDER
            value: /home/airflow/gcs/dags
          - name: GCSFUSE_EXTRACTED
            value: "TRUE"
          - name: GCP_PROJECT
            value: digraph-2021
          - name: COMPOSER_LOCATION
            value: us-central1
          - name: COMPOSER_PYTHON_VERSION
            value: "3"
          - name: COMPOSER_ENVIRONMENT
            value: airflow-bed
          - name: GKE_CLUSTER_NAME
            value: us-central1-airflow-bed-152ea37f-gke
          - name: CONTAINER_NAME
            value: airflow-webserver
          image: us-central1-docker.pkg.dev/digraph-2021/composer-images-us-central1-airflow-bed-152ea37f-gke/33e72803-04e7-4edc-ab9c-b7e19661b109
          imagePullPolicy: IfNotPresent
          livenessProbe:
            exec:
              command:
              - curl
              - --max-time
              - "30"
              - localhost:8080/_ah/health
            failureThreshold: 3
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 40
          name: airflow-webserver
          resources:
            limits:
              cpu: 700m
              ephemeral-storage: 1843Mi
              memory: 1433Mi
            requests:
              cpu: 700m
              ephemeral-storage: 1843Mi
              memory: 1433Mi
          securityContext:
            capabilities:
              drop:
              - NET_RAW
          startupProbe:
            exec:
              command:
              - curl
              - --max-time
              - "30"
              - localhost:8080/_ah/health
            failureThreshold: 18
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 40
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /etc/airflow/airflow_cfg
            name: airflow-config
          - mountPath: /home/airflow/gcs
            name: gcsdir
          - mountPath: /home/airflow/container-comms
            name: container-comms
          - mountPath: /home/airflow/gcsfuse
            mountPropagation: HostToContainer
            name: gcsfuse
        - args:
          - /home/airflow/gcs
          env:
          - name: GCS_BUCKET
            value: us-central1-airflow-bed-152ea37f-bucket
          - name: SQL_DATABASE
            value: composer-2-0-19-airflow-2-2-5-e4894b1f
          - name: SQL_USER
            value: root
          - name: SQL_PASSWORD
            valueFrom:
              secretKeyRef:
                key: sql_password
                name: airflow-secrets
          - name: COMPOSER_GKE_NAME
            value: us-central1-airflow-bed-152ea37f-gke
          - name: SKIP_SYNCING_DAGS
            value: "TRUE"
          - name: AUTOGKE
            value: "TRUE"
          - name: COMPOSER_GKE_LOCATION
            value: us-central1
          image: us-docker.pkg.dev/cloud-airflow-releaser/gcs-syncd/gcs-syncd:cloud_composer_service_2022-06-28-RC0
          imagePullPolicy: IfNotPresent
          name: gcs-syncd
          resources:
            limits:
              cpu: 150m
              ephemeral-storage: 102Mi
              memory: 307Mi
            requests:
              cpu: 150m
              ephemeral-storage: 102Mi
              memory: 307Mi
          securityContext:
            capabilities:
              drop:
              - NET_RAW
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /home/airflow/gcs
            name: gcsdir
        - args:
          - --backend=2bddbf62981044cc9b7df0b2fb317261
          - --proxy=https://us-central1.composer.cloud.google.com/tun/m/4592f092208ecc84946b8f8f8016274df1b36a14/
          - --shim-websockets=true
          - --shim-path=websocket-shim
          - --forward-user-id=true
          - --health-check-path=/_ah/health
          - --health-check-interval-seconds=5
          image: us-docker.pkg.dev/cloud-airflow-releaser/composer-inverting-proxy/composer-inverting-proxy:cloud_composer_service_2022-06-28-RC0
          imagePullPolicy: IfNotPresent
          name: agent
          resources:
            limits:
              cpu: 150m
              ephemeral-storage: 102Mi
              memory: 307Mi
            requests:
              cpu: 150m
              ephemeral-storage: 102Mi
              memory: 307Mi
          securityContext:
            capabilities:
              drop:
              - NET_RAW
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          seccompProfile:
            type: RuntimeDefault
        terminationGracePeriodSeconds: 30
        volumes:
        - configMap:
            defaultMode: 420
            name: airflow-configmap
          name: airflow-config
        - emptyDir: {}
          name: gcsdir
        - hostPath:
            path: /var/composer/gcs_mount_status
            type: ""
          name: container-comms
        - hostPath:
            path: /var/composer/gcs_mount
            type: ""
          name: gcsfuse
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-07-18T19:54:14Z"
      lastUpdateTime: "2022-09-07T22:53:26Z"
      message: ReplicaSet "airflow-webserver-5cd558768" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-26T05:10:15Z"
      lastUpdateTime: "2023-02-26T05:10:15Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 4
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "2"
    creationTimestamp: "2022-02-15T05:55:21Z"
    generation: 2
    labels:
      run: airflow-monitoring
    name: airflow-monitoring
    namespace: composer-system
    resourceVersion: "306942334"
    uid: 919e0f15-eecb-4984-bf03-577b228e9c22
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        run: airflow-monitoring
    strategy:
      rollingUpdate:
        maxSurge: 1
        maxUnavailable: 1
      type: RollingUpdate
    template:
      metadata:
        creationTimestamp: null
        labels:
          composer-system-pod: "true"
          run: airflow-monitoring
      spec:
        containers:
        - args:
          - /var/local/setup_monitoring.sh
          command:
          - /bin/bash
          - -ce
          env:
          - name: GCP_PROJECT
            value: digraph-2021
          - name: COMPOSER_LOCATION
            value: us-central1
          - name: COMPOSER_GKE_ZONE
          - name: AUTOGKE
            value: "TRUE"
          - name: COMPOSER_GKE_LOCATION
            value: us-central1
          - name: GCS_BUCKET
            value: us-central1-airflow-bed-152ea37f-bucket
          - name: COMPOSER_PYTHON_VERSION
            value: "3"
          - name: REDIS_HOST
            value: airflow-redis-service.composer-system
          - name: AIRFLOW_DATABASE_VERSION
            value: POSTGRES_13
          - name: SQL_HOST
            value: airflow-sqlproxy-service.composer-system.svc.cluster.local
          - name: SQL_DATABASE
            value: composer-2-0-19-airflow-2-2-5-e4894b1f
          - name: COMPOSER_ENVIRONMENT
            value: airflow-bed
          - name: IMAGE_VERSION
            value: composer-2-0-19-airflow-2-2-5
          - name: GKE_CLUSTER_NAME
            value: us-central1-airflow-bed-152ea37f-gke
          - name: CONTAINER_NAME
            value: airflow-monitoring
          - name: CELERY_APP_NAME
            value: airflow.executors.celery_executor
          - name: DEFAULT_QUEUE
            value: default
          - name: STATSD_PORT
            value: "8125"
          - name: REDIS_PORT
            value: "6379"
          - name: SQL_USER
            value: root
          - name: SQL_PASSWORD
            valueFrom:
              secretKeyRef:
                key: sql_password
                name: airflow-secrets-default
          - name: COMPOSER_ENVIRONMENT_TYPE
            value: PUBLIC_IP
          - name: TENANT_PROJECT
            value: a06a05266fdf96c46p-tp
          - name: SQL_INSTANCE_NAME
            value: a06a05266fdf96c46p-tp:us-central1-airflow-bed-152ea37f-sql
          - name: AUTOSCALING_ENABLED
            value: "TRUE"
          - name: VERSIONED_NAMESPACE
            value: composer-2-0-19-airflow-2-2-5-e4894b1f
          image: us-docker.pkg.dev/cloud-airflow-releaser/airflow-monitoring/airflow-monitoring:cloud_composer_service_2022-06-28-RC0
          imagePullPolicy: IfNotPresent
          livenessProbe:
            exec:
              command:
              - /home/airflow/metric_prober.py
            failureThreshold: 1
            initialDelaySeconds: 300
            periodSeconds: 180
            successThreshold: 1
            timeoutSeconds: 30
          name: airflow-monitoring
          resources:
            limits:
              cpu: 300m
              ephemeral-storage: 100Mi
              memory: 600Mi
            requests:
              cpu: 200m
              ephemeral-storage: 100Mi
              memory: 600Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        - args:
          - /home/airflow/start_task_monitor.sh
          command:
          - /bin/bash
          - -ce
          env:
          - name: COMPOSER_GKE_ZONE
          - name: COMPOSER_GKE_NAME
            value: us-central1-airflow-bed-152ea37f-gke
          - name: AUTOGKE
            value: "TRUE"
          - name: COMPOSER_GKE_LOCATION
            value: us-central1
          - name: SQL_HOST
            value: airflow-sqlproxy-service.composer-system.svc.cluster.local
          - name: SQL_USER
            value: root
          - name: SQL_PASSWORD
            valueFrom:
              secretKeyRef:
                key: sql_password
                name: airflow-secrets-default
          - name: SQL_DATABASE
            value: composer-2-0-19-airflow-2-2-5-e4894b1f
          - name: AIRFLOW_DATABASE_VERSION
            value: POSTGRES_13
          image: us-docker.pkg.dev/cloud-airflow-releaser/task-monitor/task-monitor:cloud_composer_service_2022-06-28-RC0
          imagePullPolicy: IfNotPresent
          name: task-monitor
          resources:
            limits:
              cpu: 50m
              ephemeral-storage: 100Mi
              memory: 100Mi
            requests:
              cpu: 50m
              ephemeral-storage: 100Mi
              memory: 100Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        terminationGracePeriodSeconds: 30
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-02-15T05:55:22Z"
      lastUpdateTime: "2022-02-15T05:55:22Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    - lastTransitionTime: "2022-02-15T05:55:21Z"
      lastUpdateTime: "2022-07-18T19:55:20Z"
      message: ReplicaSet "airflow-monitoring-6f87f44ddf" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    observedGeneration: 2
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "3"
    creationTimestamp: "2022-02-15T05:55:22Z"
    generation: 16
    labels:
      run: airflow-sqlproxy
    name: airflow-sqlproxy
    namespace: composer-system
    resourceVersion: "284920807"
    uid: f305126a-ffe9-4092-b2d4-47275fe0c771
  spec:
    minReadySeconds: 5
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        run: airflow-sqlproxy
    strategy:
      rollingUpdate:
        maxSurge: 50%
        maxUnavailable: 0
      type: RollingUpdate
    template:
      metadata:
        creationTimestamp: null
        labels:
          composer-system-pod: "true"
          run: airflow-sqlproxy
      spec:
        affinity:
          podAntiAffinity:
            preferredDuringSchedulingIgnoredDuringExecution:
            - podAffinityTerm:
                labelSelector:
                  matchLabels:
                    run: airflow-scheduler
                namespaces:
                - composer-2-0-19-airflow-2-2-5-e4894b1f
                topologyKey: kubernetes.io/hostname
              weight: 1
        containers:
        - env:
          - name: AIRFLOW_DATABASE_VERSION
            value: POSTGRES_13
          - name: SQL_PROXY_INSTANCES
            value: a06a05266fdf96c46p-tp:us-central1:us-central1-airflow-bed-152ea37f-sql=tcp:0.0.0.0:3306
          - name: SQL_PROXY_TERM_TIMEOUT
            value: 585s
          - name: SQL_PASSWORD
            valueFrom:
              secretKeyRef:
                key: sql_password
                name: airflow-secrets-default
          - name: HEALTHCHECK_PORT
            value: "3307"
          image: us-docker.pkg.dev/cloud-airflow-releaser/composer-cloudsql-proxy/composer-cloudsql-proxy:cloud_composer_service_2022-06-28-RC0
          imagePullPolicy: IfNotPresent
          livenessProbe:
            exec:
              command:
              - /var/local/liveness_probe.sh
            failureThreshold: 3
            initialDelaySeconds: 540
            periodSeconds: 30
            successThreshold: 1
            timeoutSeconds: 30
          name: airflow-sqlproxy
          ports:
          - containerPort: 3306
            protocol: TCP
          readinessProbe:
            exec:
              command:
              - /var/local/liveness_probe.sh
            failureThreshold: 54
            initialDelaySeconds: 10
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 30
          resources:
            limits:
              cpu: 200m
              ephemeral-storage: 100Mi
              memory: 200Mi
            requests:
              cpu: 200m
              ephemeral-storage: 100Mi
              memory: 200Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        priorityClassName: highest-priority
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        terminationGracePeriodSeconds: 600
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-02-15T05:55:22Z"
      lastUpdateTime: "2022-07-18T20:01:26Z"
      message: ReplicaSet "airflow-sqlproxy-794c8bb57b" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-26T05:13:20Z"
      lastUpdateTime: "2023-02-26T05:13:20Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 16
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "1"
    creationTimestamp: "2022-02-15T05:55:22Z"
    generation: 1
    labels:
      control-plane: worker-set-controller
    name: airflow-worker-set-controller
    namespace: composer-system
    resourceVersion: "284150870"
    uid: 84aa33d2-73a0-41e1-aaed-2bf57c54082f
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        control-plane: worker-set-controller
    strategy:
      type: Recreate
    template:
      metadata:
        annotations:
          cluster-autoscaler.kubernetes.io/safe-to-evict: "true"
        creationTimestamp: null
        labels:
          control-plane: worker-set-controller
      spec:
        containers:
        - args:
          - --metrics-addr=127.0.0.1:8080
          command:
          - /manager
          image: us-docker.pkg.dev/cloud-airflow-releaser/airflow-worker-set-controller/airflow-worker-set-controller:cloud_composer_service_2022-02-01-RC1
          imagePullPolicy: IfNotPresent
          name: manager
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        - args:
          - --secure-listen-address=0.0.0.0:8443
          - --upstream=http://127.0.0.1:8080/
          - --logtostderr=true
          - --v=10
          image: gcr.io/kubebuilder/kube-rbac-proxy:v0.5.0
          imagePullPolicy: IfNotPresent
          name: kube-rbac-proxy
          ports:
          - containerPort: 8443
            name: https
            protocol: TCP
          resources:
            limits:
              cpu: 50m
              ephemeral-storage: 100Mi
              memory: 50Mi
            requests:
              cpu: 50m
              ephemeral-storage: 100Mi
              memory: 50Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        terminationGracePeriodSeconds: 10
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-02-15T05:55:23Z"
      lastUpdateTime: "2022-02-15T05:55:29Z"
      message: ReplicaSet "airflow-worker-set-controller-f8b6c4c9b" has successfully
        progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-25T05:14:10Z"
      lastUpdateTime: "2023-02-25T05:14:10Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 1
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "3"
    creationTimestamp: "2022-09-05T13:30:32Z"
    generation: 3
    labels:
      composer-infra: critical
      run: autohealing
    name: composer-autohealing
    namespace: composer-system
    resourceVersion: "284150802"
    uid: 47ee47e1-bf20-469c-82ac-14ffe082982d
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        run: autohealing
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        creationTimestamp: null
        labels:
          composer-infra: critical
          run: autohealing
      spec:
        containers:
        - env:
          - name: COMPOSER_V2
            value: "TRUE"
          - name: AUTOHEALER_NO_DRY_RUN_ACTIONS
            value: DELETE_ANETD_POD
          - name: PROJECT_ID
            value: digraph-2021
          - name: LOCATION
            value: us-central1
          - name: ENVIRONMENT_NAME
            value: airflow-bed
          - name: ENVIRONMENT_UUID
            value: 2bddbf62-9810-44cc-9b7d-f0b2fb317261
          - name: HEARTBEAT_FILEPATH
            value: /tmp/heartbeat
          image: us-docker.pkg.dev/cloud-airflow-releaser/composer-autohealing/composer-autohealing:dnsfix_rc4
          imagePullPolicy: IfNotPresent
          livenessProbe:
            exec:
              command:
              - /opt/probe
            failureThreshold: 3
            initialDelaySeconds: 60
            periodSeconds: 180
            successThreshold: 1
            timeoutSeconds: 30
          name: composer-autohealing
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        serviceAccount: privileged
        serviceAccountName: privileged
        terminationGracePeriodSeconds: 90
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-09-05T13:30:32Z"
      lastUpdateTime: "2023-02-12T12:40:23Z"
      message: ReplicaSet "composer-autohealing-68c9f5db8b" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-25T05:14:05Z"
      lastUpdateTime: "2023-02-25T05:14:05Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 3
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "2"
      kubectl.kubernetes.io/last-applied-configuration: |
        {"apiVersion":"apps/v1","kind":"Deployment","metadata":{"annotations":{},"labels":{"k8s-app":"custom-metrics-stackdriver-adapter","run":"custom-metrics-stackdriver-adapter"},"name":"custom-metrics-stackdriver-adapter","namespace":"composer-system"},"spec":{"replicas":1,"selector":{"matchLabels":{"k8s-app":"custom-metrics-stackdriver-adapter","run":"custom-metrics-stackdriver-adapter"}},"template":{"metadata":{"annotations":{"cluster-autoscaler.kubernetes.io/safe-to-evict":"true"},"labels":{"k8s-app":"custom-metrics-stackdriver-adapter","kubernetes.io/cluster-service":"true","run":"custom-metrics-stackdriver-adapter"}},"spec":{"containers":[{"command":["/adapter","--use-new-resource-model=true"],"image":"gcr.io/gke-release/custom-metrics-stackdriver-adapter:v0.12.0-gke.0","imagePullPolicy":"Always","name":"pod-custom-metrics-stackdriver-adapter","resources":{"limits":{"cpu":"100m","ephemeral-storage":"100Mi","memory":"200Mi"},"requests":{"cpu":"50m","ephemeral-storage":"100Mi","memory":"200Mi"}}}],"serviceAccountName":"custom-metrics-stackdriver-adapter"}}}}
    creationTimestamp: "2022-02-15T05:55:15Z"
    generation: 2
    labels:
      k8s-app: custom-metrics-stackdriver-adapter
      run: custom-metrics-stackdriver-adapter
    name: custom-metrics-stackdriver-adapter
    namespace: composer-system
    resourceVersion: "284886570"
    uid: 4a825101-f166-4a2b-ada8-07cb1ad6e82e
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        k8s-app: custom-metrics-stackdriver-adapter
        run: custom-metrics-stackdriver-adapter
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        annotations:
          cluster-autoscaler.kubernetes.io/safe-to-evict: "true"
        creationTimestamp: null
        labels:
          k8s-app: custom-metrics-stackdriver-adapter
          kubernetes.io/cluster-service: "true"
          run: custom-metrics-stackdriver-adapter
      spec:
        containers:
        - command:
          - /adapter
          - --use-new-resource-model=true
          image: gcr.io/gke-release/custom-metrics-stackdriver-adapter:v0.12.2-gke.0
          imagePullPolicy: Always
          name: pod-custom-metrics-stackdriver-adapter
          resources:
            limits:
              cpu: 250m
              ephemeral-storage: 102Mi
              memory: 409Mi
            requests:
              cpu: 250m
              ephemeral-storage: 102Mi
              memory: 409Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        serviceAccount: custom-metrics-stackdriver-adapter
        serviceAccountName: custom-metrics-stackdriver-adapter
        terminationGracePeriodSeconds: 30
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-02-15T05:55:15Z"
      lastUpdateTime: "2022-06-13T10:26:52Z"
      message: ReplicaSet "custom-metrics-stackdriver-adapter-7b5f8c6fc" has successfully
        progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-26T04:10:41Z"
      lastUpdateTime: "2023-02-26T04:10:41Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 2
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "2"
    creationTimestamp: "2022-02-15T05:46:18Z"
    generation: 3
    labels:
      addonmanager.kubernetes.io/mode: Reconcile
      k8s-app: event-exporter
      kubernetes.io/cluster-service: "true"
      version: v0.3.9
    name: event-exporter-gke
    namespace: kube-system
    resourceVersion: "305781635"
    uid: 60bd485e-8676-47db-a5c4-134aa9b4b20c
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        k8s-app: event-exporter
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        annotations:
          components.gke.io/component-name: event-exporter
          components.gke.io/component-version: 1.1.0
        creationTimestamp: null
        labels:
          k8s-app: event-exporter
          version: v0.3.9
      spec:
        containers:
        - command:
          - /event-exporter
          - -sink-opts=-stackdriver-resource-model=new -endpoint=https://logging.googleapis.com
          - -prometheus-endpoint=:8080
          image: gke.gcr.io/event-exporter:v0.3.9-gke.0
          imagePullPolicy: IfNotPresent
          name: event-exporter
          resources: {}
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              drop:
              - all
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        - command:
          - /monitor
          - --stackdriver-prefix=container.googleapis.com/internal/addons
          - --api-override=https://monitoring.googleapis.com/
          - --source=event_exporter:http://localhost:8080?whitelisted=stackdriver_sink_received_entry_count,stackdriver_sink_request_count,stackdriver_sink_successfully_sent_entry_count
          - --pod-id=$(POD_NAME)
          - --namespace-id=$(POD_NAMESPACE)
          - --node-name=$(NODE_NAME)
          env:
          - name: POD_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.name
          - name: POD_NAMESPACE
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.namespace
          - name: NODE_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: spec.nodeName
          image: gke.gcr.io/prometheus-to-sd:v0.10.0-gke.0
          imagePullPolicy: IfNotPresent
          name: prometheus-to-sd-exporter
          resources: {}
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              drop:
              - all
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        nodeSelector:
          kubernetes.io/os: linux
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          runAsGroup: 1000
          runAsUser: 1000
        serviceAccount: event-exporter-sa
        serviceAccountName: event-exporter-sa
        terminationGracePeriodSeconds: 120
        tolerations:
        - key: components.gke.io/gke-managed-components
          operator: Exists
        volumes:
        - hostPath:
            path: /etc/ssl/certs
            type: Directory
          name: ssl-certs
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-02-15T05:46:18Z"
      lastUpdateTime: "2022-11-15T02:20:41Z"
      message: ReplicaSet "event-exporter-gke-f66d9f855" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-25T05:14:15Z"
      lastUpdateTime: "2023-02-25T05:14:15Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 3
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      components.gke.io/layer: addon
      credential-normal-mode: "true"
      deployment.kubernetes.io/revision: "5"
    creationTimestamp: "2022-02-15T05:46:34Z"
    generation: 82
    labels:
      addonmanager.kubernetes.io/mode: Reconcile
      k8s-app: konnectivity-agent
    name: konnectivity-agent
    namespace: kube-system
    resourceVersion: "305781720"
    uid: 5859a5ec-9ef1-4e4c-b2d9-db8de54b9da6
  spec:
    progressDeadlineSeconds: 600
    replicas: 3
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        k8s-app: konnectivity-agent
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        annotations:
          cluster-autoscaler.kubernetes.io/safe-to-evict: "true"
          components.gke.io/component-name: konnectivitynetworkproxy-combined
          components.gke.io/component-version: 1.4.10
        creationTimestamp: null
        labels:
          k8s-app: konnectivity-agent
      spec:
        containers:
        - args:
          - --logtostderr=true
          - --ca-cert=/var/run/secrets/kubernetes.io/serviceaccount/ca.crt
          - --proxy-server-host=35.225.148.2
          - --proxy-server-port=8132
          - --health-server-port=8093
          - --admin-server-port=8094
          - --sync-interval=5s
          - --sync-interval-cap=30s
          - --sync-forever=true
          - --probe-interval=5s
          - --keepalive-time=60s
          - --service-account-token-path=/var/run/secrets/tokens/konnectivity-agent-token
          - --v=3
          command:
          - /proxy-agent
          env:
          - name: POD_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.name
          - name: POD_NAMESPACE
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.namespace
          image: gke.gcr.io/proxy-agent:v0.0.31-gke.0
          imagePullPolicy: IfNotPresent
          livenessProbe:
            failureThreshold: 3
            httpGet:
              path: /healthz
              port: 8093
              scheme: HTTP
            initialDelaySeconds: 15
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 15
          name: konnectivity-agent
          ports:
          - containerPort: 8093
            name: metrics
            protocol: TCP
          resources:
            limits:
              memory: 125Mi
            requests:
              cpu: 10m
              memory: 30Mi
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              drop:
              - all
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /var/run/secrets/tokens
            name: konnectivity-agent-token
        dnsPolicy: ClusterFirst
        nodeSelector:
          beta.kubernetes.io/os: linux
        priorityClassName: system-cluster-critical
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          fsGroup: 1000
          runAsGroup: 1000
          runAsUser: 1000
        serviceAccount: konnectivity-agent
        serviceAccountName: konnectivity-agent
        terminationGracePeriodSeconds: 30
        tolerations:
        - key: CriticalAddonsOnly
          operator: Exists
        - effect: NoSchedule
          key: sandbox.gke.io/runtime
          operator: Equal
          value: gvisor
        - key: components.gke.io/gke-managed-components
          operator: Exists
        topologySpreadConstraints:
        - labelSelector:
            matchLabels:
              k8s-app: konnectivity-agent
          maxSkew: 1
          topologyKey: topology.kubernetes.io/zone
          whenUnsatisfiable: ScheduleAnyway
        - labelSelector:
            matchLabels:
              k8s-app: konnectivity-agent
          maxSkew: 1
          topologyKey: kubernetes.io/hostname
          whenUnsatisfiable: ScheduleAnyway
        volumes:
        - name: konnectivity-agent-token
          projected:
            defaultMode: 420
            sources:
            - serviceAccountToken:
                audience: system:konnectivity-server
                expirationSeconds: 3600
                path: konnectivity-agent-token
  status:
    availableReplicas: 3
    conditions:
    - lastTransitionTime: "2022-02-15T05:46:34Z"
      lastUpdateTime: "2022-09-28T09:26:49Z"
      message: ReplicaSet "konnectivity-agent-68b565957b" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-26T05:12:51Z"
      lastUpdateTime: "2023-02-26T05:12:51Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 82
    readyReplicas: 3
    replicas: 3
    updatedReplicas: 3
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      components.gke.io/layer: addon
      deployment.kubernetes.io/revision: "5"
    creationTimestamp: "2022-02-15T05:46:19Z"
    generation: 6
    labels:
      addonmanager.kubernetes.io/mode: Reconcile
      k8s-app: konnectivity-agent-autoscaler
      kubernetes.io/cluster-service: "true"
    name: konnectivity-agent-autoscaler
    namespace: kube-system
    resourceVersion: "284150214"
    uid: d7e6aea3-9991-423b-97fa-04fd187bc2e6
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        k8s-app: konnectivity-agent-autoscaler
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        annotations:
          cluster-autoscaler.kubernetes.io/safe-to-evict: "true"
          components.gke.io/component-name: konnectivitynetworkproxy-combined
          components.gke.io/component-version: 1.4.10
        creationTimestamp: null
        labels:
          k8s-app: konnectivity-agent-autoscaler
      spec:
        containers:
        - command:
          - /cluster-proportional-autoscaler
          - --namespace=kube-system
          - --configmap=konnectivity-agent-autoscaler-config
          - --target=deployment/konnectivity-agent
          - --logtostderr=true
          - --v=2
          image: gke.gcr.io/cluster-proportional-autoscaler:1.8.4-gke.1
          imagePullPolicy: IfNotPresent
          name: autoscaler
          resources:
            requests:
              cpu: 10m
              memory: 10M
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              drop:
              - all
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        nodeSelector:
          beta.kubernetes.io/os: linux
        priorityClassName: system-cluster-critical
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          runAsGroup: 1000
          runAsUser: 1000
        serviceAccount: konnectivity-agent-cpha
        serviceAccountName: konnectivity-agent-cpha
        terminationGracePeriodSeconds: 30
        tolerations:
        - key: CriticalAddonsOnly
          operator: Exists
        - key: components.gke.io/gke-managed-components
          operator: Exists
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-02-15T05:46:19Z"
      lastUpdateTime: "2022-09-28T09:26:47Z"
      message: ReplicaSet "konnectivity-agent-autoscaler-6dfb4f9cfb" has successfully
        progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-25T05:13:04Z"
      lastUpdateTime: "2023-02-25T05:13:04Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 6
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "6"
    creationTimestamp: "2022-02-15T05:46:17Z"
    generation: 9
    labels:
      addonmanager.kubernetes.io/mode: Reconcile
      k8s-app: kube-dns
      kubernetes.io/cluster-service: "true"
    name: kube-dns
    namespace: kube-system
    resourceVersion: "305781578"
    uid: 0db97a9d-06a8-4171-a75c-99c20663c8de
  spec:
    progressDeadlineSeconds: 600
    replicas: 2
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        k8s-app: kube-dns
    strategy:
      rollingUpdate:
        maxSurge: 10%
        maxUnavailable: 0
      type: RollingUpdate
    template:
      metadata:
        annotations:
          components.gke.io/component-name: kubedns
          prometheus.io/port: "10054"
          prometheus.io/scrape: "true"
          scheduler.alpha.kubernetes.io/critical-pod: ""
          seccomp.security.alpha.kubernetes.io/pod: runtime/default
        creationTimestamp: null
        labels:
          k8s-app: kube-dns
      spec:
        affinity:
          podAntiAffinity:
            preferredDuringSchedulingIgnoredDuringExecution:
            - podAffinityTerm:
                labelSelector:
                  matchExpressions:
                  - key: k8s-app
                    operator: In
                    values:
                    - kube-dns
                topologyKey: kubernetes.io/hostname
              weight: 100
        containers:
        - args:
          - --domain=cluster.local.
          - --dns-port=10053
          - --config-dir=/kube-dns-config
          - --v=2
          env:
          - name: PROMETHEUS_PORT
            value: "10055"
          image: gke.gcr.io/k8s-dns-kube-dns:1.22.12-gke.0
          imagePullPolicy: IfNotPresent
          livenessProbe:
            failureThreshold: 5
            httpGet:
              path: /healthcheck/kubedns
              port: 10054
              scheme: HTTP
            initialDelaySeconds: 60
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 5
          name: kubedns
          ports:
          - containerPort: 10053
            name: dns-local
            protocol: UDP
          - containerPort: 10053
            name: dns-tcp-local
            protocol: TCP
          - containerPort: 10055
            name: metrics
            protocol: TCP
          readinessProbe:
            failureThreshold: 3
            httpGet:
              path: /readiness
              port: 8081
              scheme: HTTP
            initialDelaySeconds: 3
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 5
          resources:
            limits:
              memory: 210Mi
            requests:
              cpu: 100m
              memory: 70Mi
          securityContext:
            allowPrivilegeEscalation: false
            readOnlyRootFilesystem: true
            runAsGroup: 1001
            runAsUser: 1001
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /kube-dns-config
            name: kube-dns-config
        - args:
          - -v=2
          - -logtostderr
          - -configDir=/etc/k8s/dns/dnsmasq-nanny
          - -restartDnsmasq=true
          - --
          - -k
          - --cache-size=1000
          - --no-negcache
          - --dns-forward-max=1500
          - --log-facility=-
          - --server=/cluster.local/127.0.0.1#10053
          - --server=/in-addr.arpa/127.0.0.1#10053
          - --server=/ip6.arpa/127.0.0.1#10053
          - --max-ttl=30
          - --max-cache-ttl=30
          image: gke.gcr.io/k8s-dns-dnsmasq-nanny:1.22.12-gke.0
          imagePullPolicy: IfNotPresent
          livenessProbe:
            failureThreshold: 5
            httpGet:
              path: /healthcheck/dnsmasq
              port: 10054
              scheme: HTTP
            initialDelaySeconds: 60
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 5
          name: dnsmasq
          ports:
          - containerPort: 53
            name: dns
            protocol: UDP
          - containerPort: 53
            name: dns-tcp
            protocol: TCP
          resources:
            requests:
              cpu: 150m
              memory: 20Mi
          securityContext:
            capabilities:
              add:
              - NET_BIND_SERVICE
              - SETGID
              drop:
              - all
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /etc/k8s/dns/dnsmasq-nanny
            name: kube-dns-config
        - args:
          - --v=2
          - --logtostderr
          - --probe=kubedns,127.0.0.1:10053,kubernetes.default.svc.cluster.local,5,SRV
          - --probe=dnsmasq,127.0.0.1:53,kubernetes.default.svc.cluster.local,5,SRV
          image: gke.gcr.io/k8s-dns-sidecar:1.22.12-gke.0
          imagePullPolicy: IfNotPresent
          livenessProbe:
            failureThreshold: 5
            httpGet:
              path: /metrics
              port: 10054
              scheme: HTTP
            initialDelaySeconds: 60
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 5
          name: sidecar
          ports:
          - containerPort: 10054
            name: metrics
            protocol: TCP
          resources:
            requests:
              cpu: 10m
              memory: 20Mi
          securityContext:
            allowPrivilegeEscalation: false
            readOnlyRootFilesystem: true
            runAsGroup: 1001
            runAsUser: 1001
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        - command:
          - /monitor
          - --source=kubedns:http://localhost:10054?whitelisted=probe_kubedns_latency_ms,probe_kubedns_errors,probe_dnsmasq_latency_ms,probe_dnsmasq_errors,dnsmasq_misses,dnsmasq_hits
          - --stackdriver-prefix=container.googleapis.com/internal/addons
          - --api-override=https://monitoring.googleapis.com/
          - --pod-id=$(POD_NAME)
          - --namespace-id=$(POD_NAMESPACE)
          - --v=2
          env:
          - name: POD_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.name
          - name: POD_NAMESPACE
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.namespace
          image: gke.gcr.io/prometheus-to-sd:v0.11.3-gke.0
          imagePullPolicy: IfNotPresent
          name: prometheus-to-sd
          resources: {}
          securityContext:
            allowPrivilegeEscalation: false
            readOnlyRootFilesystem: true
            runAsGroup: 1001
            runAsUser: 1001
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: Default
        nodeSelector:
          kubernetes.io/os: linux
        priorityClassName: system-cluster-critical
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          fsGroup: 65534
          supplementalGroups:
          - 65534
        serviceAccount: kube-dns
        serviceAccountName: kube-dns
        terminationGracePeriodSeconds: 30
        tolerations:
        - key: CriticalAddonsOnly
          operator: Exists
        - key: components.gke.io/gke-managed-components
          operator: Exists
        - effect: NoSchedule
          key: kubernetes.io/arch
          operator: Equal
          value: arm64
        volumes:
        - configMap:
            defaultMode: 420
            name: kube-dns
            optional: true
          name: kube-dns-config
  status:
    availableReplicas: 2
    conditions:
    - lastTransitionTime: "2022-02-15T05:46:18Z"
      lastUpdateTime: "2023-02-10T03:34:47Z"
      message: ReplicaSet "kube-dns-698cf6b7dc" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-26T05:13:07Z"
      lastUpdateTime: "2023-02-26T05:13:07Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 9
    readyReplicas: 2
    replicas: 2
    updatedReplicas: 2
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "2"
    creationTimestamp: "2022-02-15T05:46:17Z"
    generation: 3
    labels:
      addonmanager.kubernetes.io/mode: Reconcile
      k8s-app: kube-dns-autoscaler
      kubernetes.io/cluster-service: "true"
    name: kube-dns-autoscaler
    namespace: kube-system
    resourceVersion: "284150211"
    uid: 4fea76ed-4820-4cff-8e3b-658ff52d97cd
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        k8s-app: kube-dns-autoscaler
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        creationTimestamp: null
        labels:
          k8s-app: kube-dns-autoscaler
      spec:
        containers:
        - command:
          - /cluster-proportional-autoscaler
          - --namespace=kube-system
          - --configmap=kube-dns-autoscaler
          - --target=Deployment/kube-dns
          - --default-params={"linear":{"coresPerReplica":256,"nodesPerReplica":16,"preventSinglePointFailure":true,"includeUnschedulableNodes":true}}
          - --logtostderr=true
          - --v=2
          image: gke.gcr.io/cluster-proportional-autoscaler:1.8.4-gke.1
          imagePullPolicy: IfNotPresent
          name: autoscaler
          resources:
            requests:
              cpu: 20m
              memory: 10Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        nodeSelector:
          kubernetes.io/os: linux
        priorityClassName: system-cluster-critical
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          fsGroup: 65534
          seccompProfile:
            type: RuntimeDefault
          supplementalGroups:
          - 65534
        serviceAccount: kube-dns-autoscaler
        serviceAccountName: kube-dns-autoscaler
        terminationGracePeriodSeconds: 30
        tolerations:
        - key: CriticalAddonsOnly
          operator: Exists
        - key: components.gke.io/gke-managed-components
          operator: Exists
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-02-15T05:46:18Z"
      lastUpdateTime: "2022-07-14T16:27:53Z"
      message: ReplicaSet "kube-dns-autoscaler-f4d55555" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-25T05:13:04Z"
      lastUpdateTime: "2023-02-25T05:13:04Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 3
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      components.gke.io/layer: addon
      deployment.kubernetes.io/revision: "4"
      seccomp.security.alpha.kubernetes.io/pod: runtime/default
    creationTimestamp: "2022-02-15T05:46:19Z"
    generation: 5
    labels:
      addonmanager.kubernetes.io/mode: Reconcile
      k8s-app: glbc
      kubernetes.io/cluster-service: "true"
      kubernetes.io/name: GLBC
    name: l7-default-backend
    namespace: kube-system
    resourceVersion: "284150849"
    uid: 6cf80b6a-64f3-4228-bb77-20888278b81e
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        k8s-app: glbc
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        annotations:
          components.gke.io/component-name: l7-lb-controller-combined
          components.gke.io/component-version: 1.14.8-gke.0
          seccomp.security.alpha.kubernetes.io/pod: runtime/default
        creationTimestamp: null
        labels:
          k8s-app: glbc
          name: glbc
      spec:
        containers:
        - image: gke.gcr.io/ingress-gce-404-server-with-metrics:v1.13.4
          imagePullPolicy: IfNotPresent
          livenessProbe:
            failureThreshold: 3
            httpGet:
              path: /healthz
              port: 8080
              scheme: HTTP
            initialDelaySeconds: 30
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 5
          name: default-http-backend
          ports:
          - containerPort: 8080
            protocol: TCP
          resources:
            requests:
              cpu: 10m
              memory: 20Mi
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              drop:
              - all
            runAsGroup: 1000
            runAsUser: 1000
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        terminationGracePeriodSeconds: 30
        tolerations:
        - key: components.gke.io/gke-managed-components
          operator: Exists
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-02-15T05:46:19Z"
      lastUpdateTime: "2023-01-14T05:21:32Z"
      message: ReplicaSet "l7-default-backend-7dc577646d" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-25T05:14:09Z"
      lastUpdateTime: "2023-02-25T05:14:09Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 5
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "9"
    creationTimestamp: "2022-03-15T15:56:43Z"
    generation: 10
    labels:
      addonmanager.kubernetes.io/mode: Reconcile
      k8s-app: metrics-server
      version: v0.4.5
    name: metrics-server-v0.4.5
    namespace: kube-system
    resourceVersion: "305781911"
    uid: ff19f9a9-df5b-4ed9-9d92-674b2a598c87
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        k8s-app: metrics-server
        version: v0.4.5
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        creationTimestamp: null
        labels:
          k8s-app: metrics-server
          version: v0.4.5
        name: metrics-server
      spec:
        containers:
        - command:
          - /metrics-server
          - --metric-resolution=30s
          - --kubelet-port=10255
          - --deprecated-kubelet-completely-insecure=true
          - --kubelet-preferred-address-types=InternalIP,Hostname,InternalDNS,ExternalDNS,ExternalIP
          - --cert-dir=/tmp
          - --secure-port=10250
          image: gke.gcr.io/metrics-server-amd64:v0.4.5-gke.0
          imagePullPolicy: IfNotPresent
          name: metrics-server
          ports:
          - containerPort: 10250
            name: https
            protocol: TCP
          resources:
            limits:
              cpu: 43m
              memory: 55Mi
            requests:
              cpu: 43m
              memory: 55Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /tmp
            name: tmp-dir
        - command:
          - /pod_nanny
          - --config-dir=/etc/config
          - --cpu=40m
          - --extra-cpu=0.5m
          - --memory=35Mi
          - --extra-memory=4Mi
          - --threshold=5
          - --deployment=metrics-server-v0.4.5
          - --container=metrics-server
          - --poll-period=30000
          - --estimator=exponential
          - --scale-down-delay=24h
          - --minClusterSize=5
          - --use-metrics=true
          env:
          - name: MY_POD_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.name
          - name: MY_POD_NAMESPACE
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.namespace
          image: gke.gcr.io/addon-resizer:1.8.13-gke.0
          imagePullPolicy: IfNotPresent
          name: metrics-server-nanny
          resources:
            limits:
              cpu: 100m
              memory: 300Mi
            requests:
              cpu: 5m
              memory: 50Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /etc/config
            name: metrics-server-config-volume
        dnsPolicy: ClusterFirst
        nodeSelector:
          kubernetes.io/os: linux
        priorityClassName: system-cluster-critical
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          seccompProfile:
            type: RuntimeDefault
        serviceAccount: metrics-server
        serviceAccountName: metrics-server
        terminationGracePeriodSeconds: 30
        tolerations:
        - key: CriticalAddonsOnly
          operator: Exists
        - key: components.gke.io/gke-managed-components
          operator: Exists
        volumes:
        - configMap:
            defaultMode: 420
            name: metrics-server-config
          name: metrics-server-config-volume
        - emptyDir: {}
          name: tmp-dir
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-03-15T15:56:43Z"
      lastUpdateTime: "2022-08-06T11:19:36Z"
      message: ReplicaSet "metrics-server-v0.4.5-fb4c49dd6" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-25T05:13:16Z"
      lastUpdateTime: "2023-02-25T05:13:16Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 10
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""
apiVersion: v1
items:
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "5"
    creationTimestamp: "2022-07-18T19:54:14Z"
    generation: 5
    labels:
      run: airflow-scheduler
    name: airflow-scheduler
    namespace: composer-2-0-19-airflow-2-2-5-e4894b1f
    resourceVersion: "304279318"
    uid: 19fa28fe-2a90-40b2-9f2f-ed7b7a369a2d
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        run: airflow-scheduler
    strategy:
      rollingUpdate:
        maxSurge: 1
        maxUnavailable: 1
      type: RollingUpdate
    template:
      metadata:
        annotations:
          cluster-autoscaler.kubernetes.io/safe-to-evict: "true"
        creationTimestamp: null
        labels:
          composer-system-pod: "true"
          config_id: 34c6b681-2b2a-4726-b5b3-59b20bc435a6
          run: airflow-scheduler
      spec:
        containers:
        - args:
          - scheduler
          env:
          - name: CLOUDSDK_METRICS_ENVIRONMENT
            value: 2.2.5+composer
          - name: GCS_BUCKET
            value: us-central1-airflow-bed-152ea37f-bucket
          - name: AIRFLOW_HOME
            value: /etc/airflow
          - name: DAGS_FOLDER
            value: /home/airflow/gcs/dags
          - name: SQL_HOST
            value: airflow-sqlproxy-service.composer-system.svc.cluster.local
          - name: SQL_DATABASE
            value: composer-2-0-19-airflow-2-2-5-e4894b1f
          - name: SQL_USER
            value: root
          - name: SQL_PASSWORD
            valueFrom:
              secretKeyRef:
                key: sql_password
                name: airflow-secrets
          - name: GCSFUSE_EXTRACTED
            value: "TRUE"
          - name: COMPOSER_VERSION
            value: 2.0.19
          - name: AIRFLOW__WEBSERVER__BASE_URL
            value: https://2bddbf62981044cc9b7df0b2fb317261-dot-us-central1.composer.googleusercontent.com
          - name: AIRFLOW__CORE__SQL_ALCHEMY_CONN
            value: postgresql+psycopg2://$(SQL_USER):$(SQL_PASSWORD)@airflow-sqlproxy-service.composer-system.svc.cluster.local:3306/$(SQL_DATABASE)
          - name: AIRFLOW__CORE__FERNET_KEY
            valueFrom:
              secretKeyRef:
                key: fernet_key
                name: airflow-secrets
          - name: GCP_PROJECT
            value: digraph-2021
          - name: COMPOSER_LOCATION
            value: us-central1
          - name: COMPOSER_GKE_ZONE
          - name: COMPOSER_GKE_NAME
            value: us-central1-airflow-bed-152ea37f-gke
          - name: AUTOGKE
            value: "TRUE"
          - name: COMPOSER_GKE_LOCATION
            value: us-central1
          - name: COMPOSER_PYTHON_VERSION
            value: "3"
          - name: COMPOSER_ENVIRONMENT
            value: airflow-bed
          - name: COMPOSER_VERSIONED_NAMESPACE
            value: composer-2-0-19-airflow-2-2-5-e4894b1f
          - name: GKE_CLUSTER_NAME
            value: us-central1-airflow-bed-152ea37f-gke
          - name: POD_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.name
          - name: CONTAINER_NAME
            value: airflow-scheduler
          image: us-central1-docker.pkg.dev/digraph-2021/composer-images-us-central1-airflow-bed-152ea37f-gke/33e72803-04e7-4edc-ab9c-b7e19661b109
          imagePullPolicy: IfNotPresent
          livenessProbe:
            exec:
              command:
              - /var/local/scheduler_checker.py
            failureThreshold: 3
            initialDelaySeconds: 120
            periodSeconds: 60
            successThreshold: 1
            timeoutSeconds: 30
          name: airflow-scheduler
          resources:
            limits:
              cpu: 850m
              ephemeral-storage: 921Mi
              memory: 1632Mi
            requests:
              cpu: 850m
              ephemeral-storage: 921Mi
              memory: 1632Mi
          securityContext:
            capabilities:
              drop:
              - NET_RAW
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /etc/airflow/airflow_cfg
            name: airflow-config
          - mountPath: /home/airflow/gcs
            name: gcsdir
          - mountPath: /home/airflow/container-comms
            name: container-comms
          - mountPath: /home/airflow/gcsfuse
            mountPropagation: HostToContainer
            name: gcsfuse
        - args:
          - /home/airflow/gcs
          env:
          - name: GCS_BUCKET
            value: us-central1-airflow-bed-152ea37f-bucket
          - name: SQL_DATABASE
            value: composer-2-0-19-airflow-2-2-5-e4894b1f
          - name: SQL_USER
            value: root
          - name: SQL_PASSWORD
            valueFrom:
              secretKeyRef:
                key: sql_password
                name: airflow-secrets
          - name: COMPOSER_GKE_ZONE
          - name: COMPOSER_GKE_NAME
            value: us-central1-airflow-bed-152ea37f-gke
          - name: AUTOGKE
            value: "TRUE"
          - name: COMPOSER_GKE_LOCATION
            value: us-central1
          image: us-docker.pkg.dev/cloud-airflow-releaser/gcs-syncd/gcs-syncd:cloud_composer_service_2022-06-28-RC0
          imagePullPolicy: IfNotPresent
          name: gcs-syncd
          resources:
            limits:
              cpu: 150m
              ephemeral-storage: 102Mi
              memory: 288Mi
            requests:
              cpu: 150m
              ephemeral-storage: 102Mi
              memory: 288Mi
          securityContext:
            capabilities:
              drop:
              - NET_RAW
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /home/airflow/gcs
            name: gcsdir
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          seccompProfile:
            type: RuntimeDefault
        terminationGracePeriodSeconds: 30
        volumes:
        - configMap:
            defaultMode: 420
            name: airflow-configmap
          name: airflow-config
        - emptyDir: {}
          name: gcsdir
        - hostPath:
            path: /var/composer/gcs_mount_status
            type: ""
          name: container-comms
        - hostPath:
            path: /var/composer/gcs_mount
            type: ""
          name: gcsfuse
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-07-18T19:54:14Z"
      lastUpdateTime: "2022-07-18T19:54:14Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    - lastTransitionTime: "2022-07-18T19:54:14Z"
      lastUpdateTime: "2022-09-07T22:51:51Z"
      message: ReplicaSet "airflow-scheduler-84d7877c6d" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    observedGeneration: 5
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "4"
    creationTimestamp: "2022-07-18T19:54:14Z"
    generation: 4
    labels:
      run: airflow-webserver
    name: airflow-webserver
    namespace: composer-2-0-19-airflow-2-2-5-e4894b1f
    resourceVersion: "284918896"
    uid: c676f566-df02-4d51-8ab1-e2ed39ace0e3
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        run: airflow-webserver
    strategy:
      rollingUpdate:
        maxSurge: 1
        maxUnavailable: 0
      type: RollingUpdate
    template:
      metadata:
        creationTimestamp: null
        labels:
          config_id: 9223150f-a6e3-4a55-a110-9a3f8490f14e
          run: airflow-webserver
      spec:
        containers:
        - args:
          - webserver
          env:
          - name: AIRFLOW_WEBSERVER
            value: "True"
          - name: SQL_HOST
            value: airflow-sqlproxy-service.composer-system.svc.cluster.local
          - name: SQL_DATABASE
            value: composer-2-0-19-airflow-2-2-5-e4894b1f
          - name: SQL_USER
            value: root
          - name: SQL_PASSWORD
            valueFrom:
              secretKeyRef:
                key: sql_password
                name: airflow-secrets
          - name: COMPOSER_VERSION
            value: 2.0.19
          - name: AIRFLOW__WEBSERVER__BASE_URL
            value: https://2bddbf62981044cc9b7df0b2fb317261-dot-us-central1.composer.googleusercontent.com
          - name: AIRFLOW__WEBSERVER__JWT_PUBLIC_KEY_URL
            value: https://us-central1.composer.cloud.google.com/jwt-public-key
          - name: AIRFLOW__WEBSERVER__INVERTING_PROXY_BACKEND_ID
            value: 2bddbf62981044cc9b7df0b2fb317261
          - name: AIRFLOW__WEBSERVER__SECRET_KEY
            value: some-random-id
          - name: AIRFLOW__WEBSERVER__UPDATE_FAB_PERMS
            value: "True"
          - name: AIRFLOW__CORE__FERNET_KEY
            valueFrom:
              secretKeyRef:
                key: fernet_key
                name: airflow-secrets
          - name: AIRFLOW__CORE__SQL_ALCHEMY_CONN
            value: postgresql+psycopg2://$(SQL_USER):$(SQL_PASSWORD)@airflow-sqlproxy-service.composer-system.svc.cluster.local:3306/$(SQL_DATABASE)
          - name: CLOUDSDK_METRICS_ENVIRONMENT
            value: 2.2.5+composer
          - name: GCS_BUCKET
            value: us-central1-airflow-bed-152ea37f-bucket
          - name: AIRFLOW_HOME
            value: /etc/airflow
          - name: DAGS_FOLDER
            value: /home/airflow/gcs/dags
          - name: GCSFUSE_EXTRACTED
            value: "TRUE"
          - name: GCP_PROJECT
            value: digraph-2021
          - name: COMPOSER_LOCATION
            value: us-central1
          - name: COMPOSER_PYTHON_VERSION
            value: "3"
          - name: COMPOSER_ENVIRONMENT
            value: airflow-bed
          - name: GKE_CLUSTER_NAME
            value: us-central1-airflow-bed-152ea37f-gke
          - name: CONTAINER_NAME
            value: airflow-webserver
          image: us-central1-docker.pkg.dev/digraph-2021/composer-images-us-central1-airflow-bed-152ea37f-gke/33e72803-04e7-4edc-ab9c-b7e19661b109
          imagePullPolicy: IfNotPresent
          livenessProbe:
            exec:
              command:
              - curl
              - --max-time
              - "30"
              - localhost:8080/_ah/health
            failureThreshold: 3
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 40
          name: airflow-webserver
          resources:
            limits:
              cpu: 700m
              ephemeral-storage: 1843Mi
              memory: 1433Mi
            requests:
              cpu: 700m
              ephemeral-storage: 1843Mi
              memory: 1433Mi
          securityContext:
            capabilities:
              drop:
              - NET_RAW
          startupProbe:
            exec:
              command:
              - curl
              - --max-time
              - "30"
              - localhost:8080/_ah/health
            failureThreshold: 18
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 40
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /etc/airflow/airflow_cfg
            name: airflow-config
          - mountPath: /home/airflow/gcs
            name: gcsdir
          - mountPath: /home/airflow/container-comms
            name: container-comms
          - mountPath: /home/airflow/gcsfuse
            mountPropagation: HostToContainer
            name: gcsfuse
        - args:
          - /home/airflow/gcs
          env:
          - name: GCS_BUCKET
            value: us-central1-airflow-bed-152ea37f-bucket
          - name: SQL_DATABASE
            value: composer-2-0-19-airflow-2-2-5-e4894b1f
          - name: SQL_USER
            value: root
          - name: SQL_PASSWORD
            valueFrom:
              secretKeyRef:
                key: sql_password
                name: airflow-secrets
          - name: COMPOSER_GKE_NAME
            value: us-central1-airflow-bed-152ea37f-gke
          - name: SKIP_SYNCING_DAGS
            value: "TRUE"
          - name: AUTOGKE
            value: "TRUE"
          - name: COMPOSER_GKE_LOCATION
            value: us-central1
          image: us-docker.pkg.dev/cloud-airflow-releaser/gcs-syncd/gcs-syncd:cloud_composer_service_2022-06-28-RC0
          imagePullPolicy: IfNotPresent
          name: gcs-syncd
          resources:
            limits:
              cpu: 150m
              ephemeral-storage: 102Mi
              memory: 307Mi
            requests:
              cpu: 150m
              ephemeral-storage: 102Mi
              memory: 307Mi
          securityContext:
            capabilities:
              drop:
              - NET_RAW
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /home/airflow/gcs
            name: gcsdir
        - args:
          - --backend=2bddbf62981044cc9b7df0b2fb317261
          - --proxy=https://us-central1.composer.cloud.google.com/tun/m/4592f092208ecc84946b8f8f8016274df1b36a14/
          - --shim-websockets=true
          - --shim-path=websocket-shim
          - --forward-user-id=true
          - --health-check-path=/_ah/health
          - --health-check-interval-seconds=5
          image: us-docker.pkg.dev/cloud-airflow-releaser/composer-inverting-proxy/composer-inverting-proxy:cloud_composer_service_2022-06-28-RC0
          imagePullPolicy: IfNotPresent
          name: agent
          resources:
            limits:
              cpu: 150m
              ephemeral-storage: 102Mi
              memory: 307Mi
            requests:
              cpu: 150m
              ephemeral-storage: 102Mi
              memory: 307Mi
          securityContext:
            capabilities:
              drop:
              - NET_RAW
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          seccompProfile:
            type: RuntimeDefault
        terminationGracePeriodSeconds: 30
        volumes:
        - configMap:
            defaultMode: 420
            name: airflow-configmap
          name: airflow-config
        - emptyDir: {}
          name: gcsdir
        - hostPath:
            path: /var/composer/gcs_mount_status
            type: ""
          name: container-comms
        - hostPath:
            path: /var/composer/gcs_mount
            type: ""
          name: gcsfuse
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-07-18T19:54:14Z"
      lastUpdateTime: "2022-09-07T22:53:26Z"
      message: ReplicaSet "airflow-webserver-5cd558768" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-26T05:10:15Z"
      lastUpdateTime: "2023-02-26T05:10:15Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 4
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "2"
    creationTimestamp: "2022-02-15T05:55:21Z"
    generation: 2
    labels:
      run: airflow-monitoring
    name: airflow-monitoring
    namespace: composer-system
    resourceVersion: "310754235"
    uid: 919e0f15-eecb-4984-bf03-577b228e9c22
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        run: airflow-monitoring
    strategy:
      rollingUpdate:
        maxSurge: 1
        maxUnavailable: 1
      type: RollingUpdate
    template:
      metadata:
        creationTimestamp: null
        labels:
          composer-system-pod: "true"
          run: airflow-monitoring
      spec:
        containers:
        - args:
          - /var/local/setup_monitoring.sh
          command:
          - /bin/bash
          - -ce
          env:
          - name: GCP_PROJECT
            value: digraph-2021
          - name: COMPOSER_LOCATION
            value: us-central1
          - name: COMPOSER_GKE_ZONE
          - name: AUTOGKE
            value: "TRUE"
          - name: COMPOSER_GKE_LOCATION
            value: us-central1
          - name: GCS_BUCKET
            value: us-central1-airflow-bed-152ea37f-bucket
          - name: COMPOSER_PYTHON_VERSION
            value: "3"
          - name: REDIS_HOST
            value: airflow-redis-service.composer-system
          - name: AIRFLOW_DATABASE_VERSION
            value: POSTGRES_13
          - name: SQL_HOST
            value: airflow-sqlproxy-service.composer-system.svc.cluster.local
          - name: SQL_DATABASE
            value: composer-2-0-19-airflow-2-2-5-e4894b1f
          - name: COMPOSER_ENVIRONMENT
            value: airflow-bed
          - name: IMAGE_VERSION
            value: composer-2-0-19-airflow-2-2-5
          - name: GKE_CLUSTER_NAME
            value: us-central1-airflow-bed-152ea37f-gke
          - name: CONTAINER_NAME
            value: airflow-monitoring
          - name: CELERY_APP_NAME
            value: airflow.executors.celery_executor
          - name: DEFAULT_QUEUE
            value: default
          - name: STATSD_PORT
            value: "8125"
          - name: REDIS_PORT
            value: "6379"
          - name: SQL_USER
            value: root
          - name: SQL_PASSWORD
            valueFrom:
              secretKeyRef:
                key: sql_password
                name: airflow-secrets-default
          - name: COMPOSER_ENVIRONMENT_TYPE
            value: PUBLIC_IP
          - name: TENANT_PROJECT
            value: a06a05266fdf96c46p-tp
          - name: SQL_INSTANCE_NAME
            value: a06a05266fdf96c46p-tp:us-central1-airflow-bed-152ea37f-sql
          - name: AUTOSCALING_ENABLED
            value: "TRUE"
          - name: VERSIONED_NAMESPACE
            value: composer-2-0-19-airflow-2-2-5-e4894b1f
          image: us-docker.pkg.dev/cloud-airflow-releaser/airflow-monitoring/airflow-monitoring:cloud_composer_service_2022-06-28-RC0
          imagePullPolicy: IfNotPresent
          livenessProbe:
            exec:
              command:
              - /home/airflow/metric_prober.py
            failureThreshold: 1
            initialDelaySeconds: 300
            periodSeconds: 180
            successThreshold: 1
            timeoutSeconds: 30
          name: airflow-monitoring
          resources:
            limits:
              cpu: 300m
              ephemeral-storage: 100Mi
              memory: 600Mi
            requests:
              cpu: 200m
              ephemeral-storage: 100Mi
              memory: 600Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        - args:
          - /home/airflow/start_task_monitor.sh
          command:
          - /bin/bash
          - -ce
          env:
          - name: COMPOSER_GKE_ZONE
          - name: COMPOSER_GKE_NAME
            value: us-central1-airflow-bed-152ea37f-gke
          - name: AUTOGKE
            value: "TRUE"
          - name: COMPOSER_GKE_LOCATION
            value: us-central1
          - name: SQL_HOST
            value: airflow-sqlproxy-service.composer-system.svc.cluster.local
          - name: SQL_USER
            value: root
          - name: SQL_PASSWORD
            valueFrom:
              secretKeyRef:
                key: sql_password
                name: airflow-secrets-default
          - name: SQL_DATABASE
            value: composer-2-0-19-airflow-2-2-5-e4894b1f
          - name: AIRFLOW_DATABASE_VERSION
            value: POSTGRES_13
          image: us-docker.pkg.dev/cloud-airflow-releaser/task-monitor/task-monitor:cloud_composer_service_2022-06-28-RC0
          imagePullPolicy: IfNotPresent
          name: task-monitor
          resources:
            limits:
              cpu: 50m
              ephemeral-storage: 100Mi
              memory: 100Mi
            requests:
              cpu: 50m
              ephemeral-storage: 100Mi
              memory: 100Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        terminationGracePeriodSeconds: 30
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-02-15T05:55:22Z"
      lastUpdateTime: "2022-02-15T05:55:22Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    - lastTransitionTime: "2022-02-15T05:55:21Z"
      lastUpdateTime: "2022-07-18T19:55:20Z"
      message: ReplicaSet "airflow-monitoring-6f87f44ddf" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    observedGeneration: 2
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "3"
    creationTimestamp: "2022-02-15T05:55:22Z"
    generation: 16
    labels:
      run: airflow-sqlproxy
    name: airflow-sqlproxy
    namespace: composer-system
    resourceVersion: "284920807"
    uid: f305126a-ffe9-4092-b2d4-47275fe0c771
  spec:
    minReadySeconds: 5
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        run: airflow-sqlproxy
    strategy:
      rollingUpdate:
        maxSurge: 50%
        maxUnavailable: 0
      type: RollingUpdate
    template:
      metadata:
        creationTimestamp: null
        labels:
          composer-system-pod: "true"
          run: airflow-sqlproxy
      spec:
        affinity:
          podAntiAffinity:
            preferredDuringSchedulingIgnoredDuringExecution:
            - podAffinityTerm:
                labelSelector:
                  matchLabels:
                    run: airflow-scheduler
                namespaces:
                - composer-2-0-19-airflow-2-2-5-e4894b1f
                topologyKey: kubernetes.io/hostname
              weight: 1
        containers:
        - env:
          - name: AIRFLOW_DATABASE_VERSION
            value: POSTGRES_13
          - name: SQL_PROXY_INSTANCES
            value: a06a05266fdf96c46p-tp:us-central1:us-central1-airflow-bed-152ea37f-sql=tcp:0.0.0.0:3306
          - name: SQL_PROXY_TERM_TIMEOUT
            value: 585s
          - name: SQL_PASSWORD
            valueFrom:
              secretKeyRef:
                key: sql_password
                name: airflow-secrets-default
          - name: HEALTHCHECK_PORT
            value: "3307"
          image: us-docker.pkg.dev/cloud-airflow-releaser/composer-cloudsql-proxy/composer-cloudsql-proxy:cloud_composer_service_2022-06-28-RC0
          imagePullPolicy: IfNotPresent
          livenessProbe:
            exec:
              command:
              - /var/local/liveness_probe.sh
            failureThreshold: 3
            initialDelaySeconds: 540
            periodSeconds: 30
            successThreshold: 1
            timeoutSeconds: 30
          name: airflow-sqlproxy
          ports:
          - containerPort: 3306
            protocol: TCP
          readinessProbe:
            exec:
              command:
              - /var/local/liveness_probe.sh
            failureThreshold: 54
            initialDelaySeconds: 10
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 30
          resources:
            limits:
              cpu: 200m
              ephemeral-storage: 100Mi
              memory: 200Mi
            requests:
              cpu: 200m
              ephemeral-storage: 100Mi
              memory: 200Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        priorityClassName: highest-priority
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        terminationGracePeriodSeconds: 600
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-02-15T05:55:22Z"
      lastUpdateTime: "2022-07-18T20:01:26Z"
      message: ReplicaSet "airflow-sqlproxy-794c8bb57b" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-26T05:13:20Z"
      lastUpdateTime: "2023-02-26T05:13:20Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 16
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "1"
    creationTimestamp: "2022-02-15T05:55:22Z"
    generation: 1
    labels:
      control-plane: worker-set-controller
    name: airflow-worker-set-controller
    namespace: composer-system
    resourceVersion: "284150870"
    uid: 84aa33d2-73a0-41e1-aaed-2bf57c54082f
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        control-plane: worker-set-controller
    strategy:
      type: Recreate
    template:
      metadata:
        annotations:
          cluster-autoscaler.kubernetes.io/safe-to-evict: "true"
        creationTimestamp: null
        labels:
          control-plane: worker-set-controller
      spec:
        containers:
        - args:
          - --metrics-addr=127.0.0.1:8080
          command:
          - /manager
          image: us-docker.pkg.dev/cloud-airflow-releaser/airflow-worker-set-controller/airflow-worker-set-controller:cloud_composer_service_2022-02-01-RC1
          imagePullPolicy: IfNotPresent
          name: manager
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        - args:
          - --secure-listen-address=0.0.0.0:8443
          - --upstream=http://127.0.0.1:8080/
          - --logtostderr=true
          - --v=10
          image: gcr.io/kubebuilder/kube-rbac-proxy:v0.5.0
          imagePullPolicy: IfNotPresent
          name: kube-rbac-proxy
          ports:
          - containerPort: 8443
            name: https
            protocol: TCP
          resources:
            limits:
              cpu: 50m
              ephemeral-storage: 100Mi
              memory: 50Mi
            requests:
              cpu: 50m
              ephemeral-storage: 100Mi
              memory: 50Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        terminationGracePeriodSeconds: 10
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-02-15T05:55:23Z"
      lastUpdateTime: "2022-02-15T05:55:29Z"
      message: ReplicaSet "airflow-worker-set-controller-f8b6c4c9b" has successfully
        progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-25T05:14:10Z"
      lastUpdateTime: "2023-02-25T05:14:10Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 1
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "3"
    creationTimestamp: "2022-09-05T13:30:32Z"
    generation: 3
    labels:
      composer-infra: critical
      run: autohealing
    name: composer-autohealing
    namespace: composer-system
    resourceVersion: "284150802"
    uid: 47ee47e1-bf20-469c-82ac-14ffe082982d
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        run: autohealing
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        creationTimestamp: null
        labels:
          composer-infra: critical
          run: autohealing
      spec:
        containers:
        - env:
          - name: COMPOSER_V2
            value: "TRUE"
          - name: AUTOHEALER_NO_DRY_RUN_ACTIONS
            value: DELETE_ANETD_POD
          - name: PROJECT_ID
            value: digraph-2021
          - name: LOCATION
            value: us-central1
          - name: ENVIRONMENT_NAME
            value: airflow-bed
          - name: ENVIRONMENT_UUID
            value: 2bddbf62-9810-44cc-9b7d-f0b2fb317261
          - name: HEARTBEAT_FILEPATH
            value: /tmp/heartbeat
          image: us-docker.pkg.dev/cloud-airflow-releaser/composer-autohealing/composer-autohealing:dnsfix_rc4
          imagePullPolicy: IfNotPresent
          livenessProbe:
            exec:
              command:
              - /opt/probe
            failureThreshold: 3
            initialDelaySeconds: 60
            periodSeconds: 180
            successThreshold: 1
            timeoutSeconds: 30
          name: composer-autohealing
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        serviceAccount: privileged
        serviceAccountName: privileged
        terminationGracePeriodSeconds: 90
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-09-05T13:30:32Z"
      lastUpdateTime: "2023-02-12T12:40:23Z"
      message: ReplicaSet "composer-autohealing-68c9f5db8b" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-25T05:14:05Z"
      lastUpdateTime: "2023-02-25T05:14:05Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 3
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "2"
      kubectl.kubernetes.io/last-applied-configuration: |
        {"apiVersion":"apps/v1","kind":"Deployment","metadata":{"annotations":{},"labels":{"k8s-app":"custom-metrics-stackdriver-adapter","run":"custom-metrics-stackdriver-adapter"},"name":"custom-metrics-stackdriver-adapter","namespace":"composer-system"},"spec":{"replicas":1,"selector":{"matchLabels":{"k8s-app":"custom-metrics-stackdriver-adapter","run":"custom-metrics-stackdriver-adapter"}},"template":{"metadata":{"annotations":{"cluster-autoscaler.kubernetes.io/safe-to-evict":"true"},"labels":{"k8s-app":"custom-metrics-stackdriver-adapter","kubernetes.io/cluster-service":"true","run":"custom-metrics-stackdriver-adapter"}},"spec":{"containers":[{"command":["/adapter","--use-new-resource-model=true"],"image":"gcr.io/gke-release/custom-metrics-stackdriver-adapter:v0.12.0-gke.0","imagePullPolicy":"Always","name":"pod-custom-metrics-stackdriver-adapter","resources":{"limits":{"cpu":"100m","ephemeral-storage":"100Mi","memory":"200Mi"},"requests":{"cpu":"50m","ephemeral-storage":"100Mi","memory":"200Mi"}}}],"serviceAccountName":"custom-metrics-stackdriver-adapter"}}}}
    creationTimestamp: "2022-02-15T05:55:15Z"
    generation: 2
    labels:
      k8s-app: custom-metrics-stackdriver-adapter
      run: custom-metrics-stackdriver-adapter
    name: custom-metrics-stackdriver-adapter
    namespace: composer-system
    resourceVersion: "284886570"
    uid: 4a825101-f166-4a2b-ada8-07cb1ad6e82e
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        k8s-app: custom-metrics-stackdriver-adapter
        run: custom-metrics-stackdriver-adapter
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        annotations:
          cluster-autoscaler.kubernetes.io/safe-to-evict: "true"
        creationTimestamp: null
        labels:
          k8s-app: custom-metrics-stackdriver-adapter
          kubernetes.io/cluster-service: "true"
          run: custom-metrics-stackdriver-adapter
      spec:
        containers:
        - command:
          - /adapter
          - --use-new-resource-model=true
          image: gcr.io/gke-release/custom-metrics-stackdriver-adapter:v0.12.2-gke.0
          imagePullPolicy: Always
          name: pod-custom-metrics-stackdriver-adapter
          resources:
            limits:
              cpu: 250m
              ephemeral-storage: 102Mi
              memory: 409Mi
            requests:
              cpu: 250m
              ephemeral-storage: 102Mi
              memory: 409Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        serviceAccount: custom-metrics-stackdriver-adapter
        serviceAccountName: custom-metrics-stackdriver-adapter
        terminationGracePeriodSeconds: 30
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-02-15T05:55:15Z"
      lastUpdateTime: "2022-06-13T10:26:52Z"
      message: ReplicaSet "custom-metrics-stackdriver-adapter-7b5f8c6fc" has successfully
        progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-26T04:10:41Z"
      lastUpdateTime: "2023-02-26T04:10:41Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 2
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "2"
    creationTimestamp: "2022-02-15T05:46:18Z"
    generation: 3
    labels:
      addonmanager.kubernetes.io/mode: Reconcile
      k8s-app: event-exporter
      kubernetes.io/cluster-service: "true"
      version: v0.3.9
    name: event-exporter-gke
    namespace: kube-system
    resourceVersion: "305781635"
    uid: 60bd485e-8676-47db-a5c4-134aa9b4b20c
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        k8s-app: event-exporter
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        annotations:
          components.gke.io/component-name: event-exporter
          components.gke.io/component-version: 1.1.0
        creationTimestamp: null
        labels:
          k8s-app: event-exporter
          version: v0.3.9
      spec:
        containers:
        - command:
          - /event-exporter
          - -sink-opts=-stackdriver-resource-model=new -endpoint=https://logging.googleapis.com
          - -prometheus-endpoint=:8080
          image: gke.gcr.io/event-exporter:v0.3.9-gke.0
          imagePullPolicy: IfNotPresent
          name: event-exporter
          resources: {}
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              drop:
              - all
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        - command:
          - /monitor
          - --stackdriver-prefix=container.googleapis.com/internal/addons
          - --api-override=https://monitoring.googleapis.com/
          - --source=event_exporter:http://localhost:8080?whitelisted=stackdriver_sink_received_entry_count,stackdriver_sink_request_count,stackdriver_sink_successfully_sent_entry_count
          - --pod-id=$(POD_NAME)
          - --namespace-id=$(POD_NAMESPACE)
          - --node-name=$(NODE_NAME)
          env:
          - name: POD_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.name
          - name: POD_NAMESPACE
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.namespace
          - name: NODE_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: spec.nodeName
          image: gke.gcr.io/prometheus-to-sd:v0.10.0-gke.0
          imagePullPolicy: IfNotPresent
          name: prometheus-to-sd-exporter
          resources: {}
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              drop:
              - all
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        nodeSelector:
          kubernetes.io/os: linux
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          runAsGroup: 1000
          runAsUser: 1000
        serviceAccount: event-exporter-sa
        serviceAccountName: event-exporter-sa
        terminationGracePeriodSeconds: 120
        tolerations:
        - key: components.gke.io/gke-managed-components
          operator: Exists
        volumes:
        - hostPath:
            path: /etc/ssl/certs
            type: Directory
          name: ssl-certs
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-02-15T05:46:18Z"
      lastUpdateTime: "2022-11-15T02:20:41Z"
      message: ReplicaSet "event-exporter-gke-f66d9f855" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-25T05:14:15Z"
      lastUpdateTime: "2023-02-25T05:14:15Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 3
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      components.gke.io/layer: addon
      credential-normal-mode: "true"
      deployment.kubernetes.io/revision: "5"
    creationTimestamp: "2022-02-15T05:46:34Z"
    generation: 82
    labels:
      addonmanager.kubernetes.io/mode: Reconcile
      k8s-app: konnectivity-agent
    name: konnectivity-agent
    namespace: kube-system
    resourceVersion: "305781720"
    uid: 5859a5ec-9ef1-4e4c-b2d9-db8de54b9da6
  spec:
    progressDeadlineSeconds: 600
    replicas: 3
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        k8s-app: konnectivity-agent
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        annotations:
          cluster-autoscaler.kubernetes.io/safe-to-evict: "true"
          components.gke.io/component-name: konnectivitynetworkproxy-combined
          components.gke.io/component-version: 1.4.10
        creationTimestamp: null
        labels:
          k8s-app: konnectivity-agent
      spec:
        containers:
        - args:
          - --logtostderr=true
          - --ca-cert=/var/run/secrets/kubernetes.io/serviceaccount/ca.crt
          - --proxy-server-host=35.225.148.2
          - --proxy-server-port=8132
          - --health-server-port=8093
          - --admin-server-port=8094
          - --sync-interval=5s
          - --sync-interval-cap=30s
          - --sync-forever=true
          - --probe-interval=5s
          - --keepalive-time=60s
          - --service-account-token-path=/var/run/secrets/tokens/konnectivity-agent-token
          - --v=3
          command:
          - /proxy-agent
          env:
          - name: POD_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.name
          - name: POD_NAMESPACE
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.namespace
          image: gke.gcr.io/proxy-agent:v0.0.31-gke.0
          imagePullPolicy: IfNotPresent
          livenessProbe:
            failureThreshold: 3
            httpGet:
              path: /healthz
              port: 8093
              scheme: HTTP
            initialDelaySeconds: 15
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 15
          name: konnectivity-agent
          ports:
          - containerPort: 8093
            name: metrics
            protocol: TCP
          resources:
            limits:
              memory: 125Mi
            requests:
              cpu: 10m
              memory: 30Mi
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              drop:
              - all
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /var/run/secrets/tokens
            name: konnectivity-agent-token
        dnsPolicy: ClusterFirst
        nodeSelector:
          beta.kubernetes.io/os: linux
        priorityClassName: system-cluster-critical
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          fsGroup: 1000
          runAsGroup: 1000
          runAsUser: 1000
        serviceAccount: konnectivity-agent
        serviceAccountName: konnectivity-agent
        terminationGracePeriodSeconds: 30
        tolerations:
        - key: CriticalAddonsOnly
          operator: Exists
        - effect: NoSchedule
          key: sandbox.gke.io/runtime
          operator: Equal
          value: gvisor
        - key: components.gke.io/gke-managed-components
          operator: Exists
        topologySpreadConstraints:
        - labelSelector:
            matchLabels:
              k8s-app: konnectivity-agent
          maxSkew: 1
          topologyKey: topology.kubernetes.io/zone
          whenUnsatisfiable: ScheduleAnyway
        - labelSelector:
            matchLabels:
              k8s-app: konnectivity-agent
          maxSkew: 1
          topologyKey: kubernetes.io/hostname
          whenUnsatisfiable: ScheduleAnyway
        volumes:
        - name: konnectivity-agent-token
          projected:
            defaultMode: 420
            sources:
            - serviceAccountToken:
                audience: system:konnectivity-server
                expirationSeconds: 3600
                path: konnectivity-agent-token
  status:
    availableReplicas: 3
    conditions:
    - lastTransitionTime: "2022-02-15T05:46:34Z"
      lastUpdateTime: "2022-09-28T09:26:49Z"
      message: ReplicaSet "konnectivity-agent-68b565957b" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-26T05:12:51Z"
      lastUpdateTime: "2023-02-26T05:12:51Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 82
    readyReplicas: 3
    replicas: 3
    updatedReplicas: 3
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      components.gke.io/layer: addon
      deployment.kubernetes.io/revision: "5"
    creationTimestamp: "2022-02-15T05:46:19Z"
    generation: 6
    labels:
      addonmanager.kubernetes.io/mode: Reconcile
      k8s-app: konnectivity-agent-autoscaler
      kubernetes.io/cluster-service: "true"
    name: konnectivity-agent-autoscaler
    namespace: kube-system
    resourceVersion: "284150214"
    uid: d7e6aea3-9991-423b-97fa-04fd187bc2e6
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        k8s-app: konnectivity-agent-autoscaler
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        annotations:
          cluster-autoscaler.kubernetes.io/safe-to-evict: "true"
          components.gke.io/component-name: konnectivitynetworkproxy-combined
          components.gke.io/component-version: 1.4.10
        creationTimestamp: null
        labels:
          k8s-app: konnectivity-agent-autoscaler
      spec:
        containers:
        - command:
          - /cluster-proportional-autoscaler
          - --namespace=kube-system
          - --configmap=konnectivity-agent-autoscaler-config
          - --target=deployment/konnectivity-agent
          - --logtostderr=true
          - --v=2
          image: gke.gcr.io/cluster-proportional-autoscaler:1.8.4-gke.1
          imagePullPolicy: IfNotPresent
          name: autoscaler
          resources:
            requests:
              cpu: 10m
              memory: 10M
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              drop:
              - all
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        nodeSelector:
          beta.kubernetes.io/os: linux
        priorityClassName: system-cluster-critical
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          runAsGroup: 1000
          runAsUser: 1000
        serviceAccount: konnectivity-agent-cpha
        serviceAccountName: konnectivity-agent-cpha
        terminationGracePeriodSeconds: 30
        tolerations:
        - key: CriticalAddonsOnly
          operator: Exists
        - key: components.gke.io/gke-managed-components
          operator: Exists
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-02-15T05:46:19Z"
      lastUpdateTime: "2022-09-28T09:26:47Z"
      message: ReplicaSet "konnectivity-agent-autoscaler-6dfb4f9cfb" has successfully
        progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-25T05:13:04Z"
      lastUpdateTime: "2023-02-25T05:13:04Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 6
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "6"
    creationTimestamp: "2022-02-15T05:46:17Z"
    generation: 9
    labels:
      addonmanager.kubernetes.io/mode: Reconcile
      k8s-app: kube-dns
      kubernetes.io/cluster-service: "true"
    name: kube-dns
    namespace: kube-system
    resourceVersion: "305781578"
    uid: 0db97a9d-06a8-4171-a75c-99c20663c8de
  spec:
    progressDeadlineSeconds: 600
    replicas: 2
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        k8s-app: kube-dns
    strategy:
      rollingUpdate:
        maxSurge: 10%
        maxUnavailable: 0
      type: RollingUpdate
    template:
      metadata:
        annotations:
          components.gke.io/component-name: kubedns
          prometheus.io/port: "10054"
          prometheus.io/scrape: "true"
          scheduler.alpha.kubernetes.io/critical-pod: ""
          seccomp.security.alpha.kubernetes.io/pod: runtime/default
        creationTimestamp: null
        labels:
          k8s-app: kube-dns
      spec:
        affinity:
          podAntiAffinity:
            preferredDuringSchedulingIgnoredDuringExecution:
            - podAffinityTerm:
                labelSelector:
                  matchExpressions:
                  - key: k8s-app
                    operator: In
                    values:
                    - kube-dns
                topologyKey: kubernetes.io/hostname
              weight: 100
        containers:
        - args:
          - --domain=cluster.local.
          - --dns-port=10053
          - --config-dir=/kube-dns-config
          - --v=2
          env:
          - name: PROMETHEUS_PORT
            value: "10055"
          image: gke.gcr.io/k8s-dns-kube-dns:1.22.12-gke.0
          imagePullPolicy: IfNotPresent
          livenessProbe:
            failureThreshold: 5
            httpGet:
              path: /healthcheck/kubedns
              port: 10054
              scheme: HTTP
            initialDelaySeconds: 60
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 5
          name: kubedns
          ports:
          - containerPort: 10053
            name: dns-local
            protocol: UDP
          - containerPort: 10053
            name: dns-tcp-local
            protocol: TCP
          - containerPort: 10055
            name: metrics
            protocol: TCP
          readinessProbe:
            failureThreshold: 3
            httpGet:
              path: /readiness
              port: 8081
              scheme: HTTP
            initialDelaySeconds: 3
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 5
          resources:
            limits:
              memory: 210Mi
            requests:
              cpu: 100m
              memory: 70Mi
          securityContext:
            allowPrivilegeEscalation: false
            readOnlyRootFilesystem: true
            runAsGroup: 1001
            runAsUser: 1001
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /kube-dns-config
            name: kube-dns-config
        - args:
          - -v=2
          - -logtostderr
          - -configDir=/etc/k8s/dns/dnsmasq-nanny
          - -restartDnsmasq=true
          - --
          - -k
          - --cache-size=1000
          - --no-negcache
          - --dns-forward-max=1500
          - --log-facility=-
          - --server=/cluster.local/127.0.0.1#10053
          - --server=/in-addr.arpa/127.0.0.1#10053
          - --server=/ip6.arpa/127.0.0.1#10053
          - --max-ttl=30
          - --max-cache-ttl=30
          image: gke.gcr.io/k8s-dns-dnsmasq-nanny:1.22.12-gke.0
          imagePullPolicy: IfNotPresent
          livenessProbe:
            failureThreshold: 5
            httpGet:
              path: /healthcheck/dnsmasq
              port: 10054
              scheme: HTTP
            initialDelaySeconds: 60
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 5
          name: dnsmasq
          ports:
          - containerPort: 53
            name: dns
            protocol: UDP
          - containerPort: 53
            name: dns-tcp
            protocol: TCP
          resources:
            requests:
              cpu: 150m
              memory: 20Mi
          securityContext:
            capabilities:
              add:
              - NET_BIND_SERVICE
              - SETGID
              drop:
              - all
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /etc/k8s/dns/dnsmasq-nanny
            name: kube-dns-config
        - args:
          - --v=2
          - --logtostderr
          - --probe=kubedns,127.0.0.1:10053,kubernetes.default.svc.cluster.local,5,SRV
          - --probe=dnsmasq,127.0.0.1:53,kubernetes.default.svc.cluster.local,5,SRV
          image: gke.gcr.io/k8s-dns-sidecar:1.22.12-gke.0
          imagePullPolicy: IfNotPresent
          livenessProbe:
            failureThreshold: 5
            httpGet:
              path: /metrics
              port: 10054
              scheme: HTTP
            initialDelaySeconds: 60
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 5
          name: sidecar
          ports:
          - containerPort: 10054
            name: metrics
            protocol: TCP
          resources:
            requests:
              cpu: 10m
              memory: 20Mi
          securityContext:
            allowPrivilegeEscalation: false
            readOnlyRootFilesystem: true
            runAsGroup: 1001
            runAsUser: 1001
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        - command:
          - /monitor
          - --source=kubedns:http://localhost:10054?whitelisted=probe_kubedns_latency_ms,probe_kubedns_errors,probe_dnsmasq_latency_ms,probe_dnsmasq_errors,dnsmasq_misses,dnsmasq_hits
          - --stackdriver-prefix=container.googleapis.com/internal/addons
          - --api-override=https://monitoring.googleapis.com/
          - --pod-id=$(POD_NAME)
          - --namespace-id=$(POD_NAMESPACE)
          - --v=2
          env:
          - name: POD_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.name
          - name: POD_NAMESPACE
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.namespace
          image: gke.gcr.io/prometheus-to-sd:v0.11.3-gke.0
          imagePullPolicy: IfNotPresent
          name: prometheus-to-sd
          resources: {}
          securityContext:
            allowPrivilegeEscalation: false
            readOnlyRootFilesystem: true
            runAsGroup: 1001
            runAsUser: 1001
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: Default
        nodeSelector:
          kubernetes.io/os: linux
        priorityClassName: system-cluster-critical
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          fsGroup: 65534
          supplementalGroups:
          - 65534
        serviceAccount: kube-dns
        serviceAccountName: kube-dns
        terminationGracePeriodSeconds: 30
        tolerations:
        - key: CriticalAddonsOnly
          operator: Exists
        - key: components.gke.io/gke-managed-components
          operator: Exists
        - effect: NoSchedule
          key: kubernetes.io/arch
          operator: Equal
          value: arm64
        volumes:
        - configMap:
            defaultMode: 420
            name: kube-dns
            optional: true
          name: kube-dns-config
  status:
    availableReplicas: 2
    conditions:
    - lastTransitionTime: "2022-02-15T05:46:18Z"
      lastUpdateTime: "2023-02-10T03:34:47Z"
      message: ReplicaSet "kube-dns-698cf6b7dc" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-26T05:13:07Z"
      lastUpdateTime: "2023-02-26T05:13:07Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 9
    readyReplicas: 2
    replicas: 2
    updatedReplicas: 2
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "2"
    creationTimestamp: "2022-02-15T05:46:17Z"
    generation: 3
    labels:
      addonmanager.kubernetes.io/mode: Reconcile
      k8s-app: kube-dns-autoscaler
      kubernetes.io/cluster-service: "true"
    name: kube-dns-autoscaler
    namespace: kube-system
    resourceVersion: "284150211"
    uid: 4fea76ed-4820-4cff-8e3b-658ff52d97cd
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        k8s-app: kube-dns-autoscaler
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        creationTimestamp: null
        labels:
          k8s-app: kube-dns-autoscaler
      spec:
        containers:
        - command:
          - /cluster-proportional-autoscaler
          - --namespace=kube-system
          - --configmap=kube-dns-autoscaler
          - --target=Deployment/kube-dns
          - --default-params={"linear":{"coresPerReplica":256,"nodesPerReplica":16,"preventSinglePointFailure":true,"includeUnschedulableNodes":true}}
          - --logtostderr=true
          - --v=2
          image: gke.gcr.io/cluster-proportional-autoscaler:1.8.4-gke.1
          imagePullPolicy: IfNotPresent
          name: autoscaler
          resources:
            requests:
              cpu: 20m
              memory: 10Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        nodeSelector:
          kubernetes.io/os: linux
        priorityClassName: system-cluster-critical
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          fsGroup: 65534
          seccompProfile:
            type: RuntimeDefault
          supplementalGroups:
          - 65534
        serviceAccount: kube-dns-autoscaler
        serviceAccountName: kube-dns-autoscaler
        terminationGracePeriodSeconds: 30
        tolerations:
        - key: CriticalAddonsOnly
          operator: Exists
        - key: components.gke.io/gke-managed-components
          operator: Exists
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-02-15T05:46:18Z"
      lastUpdateTime: "2022-07-14T16:27:53Z"
      message: ReplicaSet "kube-dns-autoscaler-f4d55555" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-25T05:13:04Z"
      lastUpdateTime: "2023-02-25T05:13:04Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 3
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      components.gke.io/layer: addon
      deployment.kubernetes.io/revision: "4"
      seccomp.security.alpha.kubernetes.io/pod: runtime/default
    creationTimestamp: "2022-02-15T05:46:19Z"
    generation: 5
    labels:
      addonmanager.kubernetes.io/mode: Reconcile
      k8s-app: glbc
      kubernetes.io/cluster-service: "true"
      kubernetes.io/name: GLBC
    name: l7-default-backend
    namespace: kube-system
    resourceVersion: "284150849"
    uid: 6cf80b6a-64f3-4228-bb77-20888278b81e
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        k8s-app: glbc
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        annotations:
          components.gke.io/component-name: l7-lb-controller-combined
          components.gke.io/component-version: 1.14.8-gke.0
          seccomp.security.alpha.kubernetes.io/pod: runtime/default
        creationTimestamp: null
        labels:
          k8s-app: glbc
          name: glbc
      spec:
        containers:
        - image: gke.gcr.io/ingress-gce-404-server-with-metrics:v1.13.4
          imagePullPolicy: IfNotPresent
          livenessProbe:
            failureThreshold: 3
            httpGet:
              path: /healthz
              port: 8080
              scheme: HTTP
            initialDelaySeconds: 30
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 5
          name: default-http-backend
          ports:
          - containerPort: 8080
            protocol: TCP
          resources:
            requests:
              cpu: 10m
              memory: 20Mi
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              drop:
              - all
            runAsGroup: 1000
            runAsUser: 1000
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
        dnsPolicy: ClusterFirst
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        terminationGracePeriodSeconds: 30
        tolerations:
        - key: components.gke.io/gke-managed-components
          operator: Exists
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-02-15T05:46:19Z"
      lastUpdateTime: "2023-01-14T05:21:32Z"
      message: ReplicaSet "l7-default-backend-7dc577646d" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-25T05:14:09Z"
      lastUpdateTime: "2023-02-25T05:14:09Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 5
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      deployment.kubernetes.io/revision: "9"
    creationTimestamp: "2022-03-15T15:56:43Z"
    generation: 10
    labels:
      addonmanager.kubernetes.io/mode: Reconcile
      k8s-app: metrics-server
      version: v0.4.5
    name: metrics-server-v0.4.5
    namespace: kube-system
    resourceVersion: "305781911"
    uid: ff19f9a9-df5b-4ed9-9d92-674b2a598c87
  spec:
    progressDeadlineSeconds: 600
    replicas: 1
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        k8s-app: metrics-server
        version: v0.4.5
    strategy:
      rollingUpdate:
        maxSurge: 25%
        maxUnavailable: 25%
      type: RollingUpdate
    template:
      metadata:
        creationTimestamp: null
        labels:
          k8s-app: metrics-server
          version: v0.4.5
        name: metrics-server
      spec:
        containers:
        - command:
          - /metrics-server
          - --metric-resolution=30s
          - --kubelet-port=10255
          - --deprecated-kubelet-completely-insecure=true
          - --kubelet-preferred-address-types=InternalIP,Hostname,InternalDNS,ExternalDNS,ExternalIP
          - --cert-dir=/tmp
          - --secure-port=10250
          image: gke.gcr.io/metrics-server-amd64:v0.4.5-gke.0
          imagePullPolicy: IfNotPresent
          name: metrics-server
          ports:
          - containerPort: 10250
            name: https
            protocol: TCP
          resources:
            limits:
              cpu: 43m
              memory: 55Mi
            requests:
              cpu: 43m
              memory: 55Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /tmp
            name: tmp-dir
        - command:
          - /pod_nanny
          - --config-dir=/etc/config
          - --cpu=40m
          - --extra-cpu=0.5m
          - --memory=35Mi
          - --extra-memory=4Mi
          - --threshold=5
          - --deployment=metrics-server-v0.4.5
          - --container=metrics-server
          - --poll-period=30000
          - --estimator=exponential
          - --scale-down-delay=24h
          - --minClusterSize=5
          - --use-metrics=true
          env:
          - name: MY_POD_NAME
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.name
          - name: MY_POD_NAMESPACE
            valueFrom:
              fieldRef:
                apiVersion: v1
                fieldPath: metadata.namespace
          image: gke.gcr.io/addon-resizer:1.8.13-gke.0
          imagePullPolicy: IfNotPresent
          name: metrics-server-nanny
          resources:
            limits:
              cpu: 100m
              memory: 300Mi
            requests:
              cpu: 5m
              memory: 50Mi
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /etc/config
            name: metrics-server-config-volume
        dnsPolicy: ClusterFirst
        nodeSelector:
          kubernetes.io/os: linux
        priorityClassName: system-cluster-critical
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext:
          seccompProfile:
            type: RuntimeDefault
        serviceAccount: metrics-server
        serviceAccountName: metrics-server
        terminationGracePeriodSeconds: 30
        tolerations:
        - key: CriticalAddonsOnly
          operator: Exists
        - key: components.gke.io/gke-managed-components
          operator: Exists
        volumes:
        - configMap:
            defaultMode: 420
            name: metrics-server-config
          name: metrics-server-config-volume
        - emptyDir: {}
          name: tmp-dir
  status:
    availableReplicas: 1
    conditions:
    - lastTransitionTime: "2022-03-15T15:56:43Z"
      lastUpdateTime: "2022-08-06T11:19:36Z"
      message: ReplicaSet "metrics-server-v0.4.5-fb4c49dd6" has successfully progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
    - lastTransitionTime: "2023-02-25T05:13:16Z"
      lastUpdateTime: "2023-02-25T05:13:16Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    observedGeneration: 10
    readyReplicas: 1
    replicas: 1
    updatedReplicas: 1
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""
`

var PARSED_LIST_JSON = `[
  {
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
      "annotations": {
        "deployment.kubernetes.io/revision": "5"
      },
      "creationTimestamp": "2022-07-18T19:54:14Z",
      "generation": 5,
      "labels": {
        "run": "airflow-scheduler"
      },
      "name": "airflow-scheduler",
      "namespace": "composer-2-0-19-airflow-2-2-5-e4894b1f",
      "resourceVersion": "304279318",
      "uid": "19fa28fe-2a90-40b2-9f2f-ed7b7a369a2d"
    },
    "spec": {
      "progressDeadlineSeconds": 600,
      "replicas": 1,
      "revisionHistoryLimit": 10,
      "selector": {
        "matchLabels": {
          "run": "airflow-scheduler"
        }
      },
      "strategy": {
        "rollingUpdate": {
          "maxSurge": 1,
          "maxUnavailable": 1
        },
        "type": "RollingUpdate"
      },
      "template": {
        "metadata": {
          "annotations": {
            "cluster-autoscaler.kubernetes.io/safe-to-evict": "true"
          },
          "creationTimestamp": null,
          "labels": {
            "composer-system-pod": "true",
            "config_id": "34c6b681-2b2a-4726-b5b3-59b20bc435a6",
            "run": "airflow-scheduler"
          }
        },
        "spec": {
          "containers": [
            {
              "args": [
                "scheduler"
              ],
              "env": [
                {
                  "name": "CLOUDSDK_METRICS_ENVIRONMENT",
                  "value": "2.2.5+composer"
                },
                {
                  "name": "GCS_BUCKET",
                  "value": "us-central1-airflow-bed-152ea37f-bucket"
                },
                {
                  "name": "AIRFLOW_HOME",
                  "value": "/etc/airflow"
                },
                {
                  "name": "DAGS_FOLDER",
                  "value": "/home/airflow/gcs/dags"
                },
                {
                  "name": "SQL_HOST",
                  "value": "airflow-sqlproxy-service.composer-system.svc.cluster.local"
                },
                {
                  "name": "SQL_DATABASE",
                  "value": "composer-2-0-19-airflow-2-2-5-e4894b1f"
                },
                {
                  "name": "SQL_USER",
                  "value": "root"
                },
                {
                  "name": "SQL_PASSWORD",
                  "valueFrom": {
                    "secretKeyRef": {
                      "key": "sql_password",
                      "name": "airflow-secrets"
                    }
                  }
                },
                {
                  "name": "GCSFUSE_EXTRACTED",
                  "value": "TRUE"
                },
                {
                  "name": "COMPOSER_VERSION",
                  "value": "2.0.19"
                },
                {
                  "name": "AIRFLOW__WEBSERVER__BASE_URL",
                  "value": "https://2bddbf62981044cc9b7df0b2fb317261-dot-us-central1.composer.googleusercontent.com"
                },
                {
                  "name": "AIRFLOW__CORE__SQL_ALCHEMY_CONN",
                  "value": "postgresql+psycopg2://$(SQL_USER):$(SQL_PASSWORD)@airflow-sqlproxy-service.composer-system.svc.cluster.local:3306/$(SQL_DATABASE)"
                },
                {
                  "name": "AIRFLOW__CORE__FERNET_KEY",
                  "valueFrom": {
                    "secretKeyRef": {
                      "key": "fernet_key",
                      "name": "airflow-secrets"
                    }
                  }
                },
                {
                  "name": "GCP_PROJECT",
                  "value": "digraph-2021"
                },
                {
                  "name": "COMPOSER_LOCATION",
                  "value": "us-central1"
                },
                {
                  "name": "COMPOSER_GKE_ZONE"
                },
                {
                  "name": "COMPOSER_GKE_NAME",
                  "value": "us-central1-airflow-bed-152ea37f-gke"
                },
                {
                  "name": "AUTOGKE",
                  "value": "TRUE"
                },
                {
                  "name": "COMPOSER_GKE_LOCATION",
                  "value": "us-central1"
                },
                {
                  "name": "COMPOSER_PYTHON_VERSION",
                  "value": "3"
                },
                {
                  "name": "COMPOSER_ENVIRONMENT",
                  "value": "airflow-bed"
                },
                {
                  "name": "COMPOSER_VERSIONED_NAMESPACE",
                  "value": "composer-2-0-19-airflow-2-2-5-e4894b1f"
                },
                {
                  "name": "GKE_CLUSTER_NAME",
                  "value": "us-central1-airflow-bed-152ea37f-gke"
                },
                {
                  "name": "POD_NAME",
                  "valueFrom": {
                    "fieldRef": {
                      "apiVersion": "v1",
                      "fieldPath": "metadata.name"
                    }
                  }
                },
                {
                  "name": "CONTAINER_NAME",
                  "value": "airflow-scheduler"
                }
              ],
              "image": "us-central1-docker.pkg.dev/digraph-2021/composer-images-us-central1-airflow-bed-152ea37f-gke/33e72803-04e7-4edc-ab9c-b7e19661b109",
              "imagePullPolicy": "IfNotPresent",
              "livenessProbe": {
                "exec": {
                  "command": [
                    "/var/local/scheduler_checker.py"
                  ]
                },
                "failureThreshold": 3,
                "initialDelaySeconds": 120,
                "periodSeconds": 60,
                "successThreshold": 1,
                "timeoutSeconds": 30
              },
              "name": "airflow-scheduler",
              "resources": {
                "limits": {
                  "cpu": "850m",
                  "ephemeral-storage": "921Mi",
                  "memory": "1632Mi"
                },
                "requests": {
                  "cpu": "850m",
                  "ephemeral-storage": "921Mi",
                  "memory": "1632Mi"
                }
              },
              "securityContext": {
                "capabilities": {
                  "drop": [
                    "NET_RAW"
                  ]
                }
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File",
              "volumeMounts": [
                {
                  "mountPath": "/etc/airflow/airflow_cfg",
                  "name": "airflow-config"
                },
                {
                  "mountPath": "/home/airflow/gcs",
                  "name": "gcsdir"
                },
                {
                  "mountPath": "/home/airflow/container-comms",
                  "name": "container-comms"
                },
                {
                  "mountPath": "/home/airflow/gcsfuse",
                  "mountPropagation": "HostToContainer",
                  "name": "gcsfuse"
                }
              ]
            },
            {
              "args": [
                "/home/airflow/gcs"
              ],
              "env": [
                {
                  "name": "GCS_BUCKET",
                  "value": "us-central1-airflow-bed-152ea37f-bucket"
                },
                {
                  "name": "SQL_DATABASE",
                  "value": "composer-2-0-19-airflow-2-2-5-e4894b1f"
                },
                {
                  "name": "SQL_USER",
                  "value": "root"
                },
                {
                  "name": "SQL_PASSWORD",
                  "valueFrom": {
                    "secretKeyRef": {
                      "key": "sql_password",
                      "name": "airflow-secrets"
                    }
                  }
                },
                {
                  "name": "COMPOSER_GKE_ZONE"
                },
                {
                  "name": "COMPOSER_GKE_NAME",
                  "value": "us-central1-airflow-bed-152ea37f-gke"
                },
                {
                  "name": "AUTOGKE",
                  "value": "TRUE"
                },
                {
                  "name": "COMPOSER_GKE_LOCATION",
                  "value": "us-central1"
                }
              ],
              "image": "us-docker.pkg.dev/cloud-airflow-releaser/gcs-syncd/gcs-syncd:cloud_composer_service_2022-06-28-RC0",
              "imagePullPolicy": "IfNotPresent",
              "name": "gcs-syncd",
              "resources": {
                "limits": {
                  "cpu": "150m",
                  "ephemeral-storage": "102Mi",
                  "memory": "288Mi"
                },
                "requests": {
                  "cpu": "150m",
                  "ephemeral-storage": "102Mi",
                  "memory": "288Mi"
                }
              },
              "securityContext": {
                "capabilities": {
                  "drop": [
                    "NET_RAW"
                  ]
                }
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File",
              "volumeMounts": [
                {
                  "mountPath": "/home/airflow/gcs",
                  "name": "gcsdir"
                }
              ]
            }
          ],
          "dnsPolicy": "ClusterFirst",
          "restartPolicy": "Always",
          "schedulerName": "default-scheduler",
          "securityContext": {
            "seccompProfile": {
              "type": "RuntimeDefault"
            }
          },
          "terminationGracePeriodSeconds": 30,
          "volumes": [
            {
              "configMap": {
                "defaultMode": 420,
                "name": "airflow-configmap"
              },
              "name": "airflow-config"
            },
            {
              "emptyDir": {},
              "name": "gcsdir"
            },
            {
              "hostPath": {
                "path": "/var/composer/gcs_mount_status",
                "type": ""
              },
              "name": "container-comms"
            },
            {
              "hostPath": {
                "path": "/var/composer/gcs_mount",
                "type": ""
              },
              "name": "gcsfuse"
            }
          ]
        }
      }
    },
    "status": {
      "availableReplicas": 1,
      "conditions": [
        {
          "lastTransitionTime": "2022-07-18T19:54:14Z",
          "lastUpdateTime": "2022-07-18T19:54:14Z",
          "message": "Deployment has minimum availability.",
          "reason": "MinimumReplicasAvailable",
          "status": "True",
          "type": "Available"
        },
        {
          "lastTransitionTime": "2022-07-18T19:54:14Z",
          "lastUpdateTime": "2022-09-07T22:51:51Z",
          "message": "ReplicaSet \"airflow-scheduler-84d7877c6d\" has successfully progressed.",
          "reason": "NewReplicaSetAvailable",
          "status": "True",
          "type": "Progressing"
        }
      ],
      "observedGeneration": 5,
      "readyReplicas": 1,
      "replicas": 1,
      "updatedReplicas": 1
    }
  },
  {
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
      "annotations": {
        "deployment.kubernetes.io/revision": "4"
      },
      "creationTimestamp": "2022-07-18T19:54:14Z",
      "generation": 4,
      "labels": {
        "run": "airflow-webserver"
      },
      "name": "airflow-webserver",
      "namespace": "composer-2-0-19-airflow-2-2-5-e4894b1f",
      "resourceVersion": "284918896",
      "uid": "c676f566-df02-4d51-8ab1-e2ed39ace0e3"
    },
    "spec": {
      "progressDeadlineSeconds": 600,
      "replicas": 1,
      "revisionHistoryLimit": 10,
      "selector": {
        "matchLabels": {
          "run": "airflow-webserver"
        }
      },
      "strategy": {
        "rollingUpdate": {
          "maxSurge": 1,
          "maxUnavailable": 0
        },
        "type": "RollingUpdate"
      },
      "template": {
        "metadata": {
          "creationTimestamp": null,
          "labels": {
            "config_id": "9223150f-a6e3-4a55-a110-9a3f8490f14e",
            "run": "airflow-webserver"
          }
        },
        "spec": {
          "containers": [
            {
              "args": [
                "webserver"
              ],
              "env": [
                {
                  "name": "AIRFLOW_WEBSERVER",
                  "value": "True"
                },
                {
                  "name": "SQL_HOST",
                  "value": "airflow-sqlproxy-service.composer-system.svc.cluster.local"
                },
                {
                  "name": "SQL_DATABASE",
                  "value": "composer-2-0-19-airflow-2-2-5-e4894b1f"
                },
                {
                  "name": "SQL_USER",
                  "value": "root"
                },
                {
                  "name": "SQL_PASSWORD",
                  "valueFrom": {
                    "secretKeyRef": {
                      "key": "sql_password",
                      "name": "airflow-secrets"
                    }
                  }
                },
                {
                  "name": "COMPOSER_VERSION",
                  "value": "2.0.19"
                },
                {
                  "name": "AIRFLOW__WEBSERVER__BASE_URL",
                  "value": "https://2bddbf62981044cc9b7df0b2fb317261-dot-us-central1.composer.googleusercontent.com"
                },
                {
                  "name": "AIRFLOW__WEBSERVER__JWT_PUBLIC_KEY_URL",
                  "value": "https://us-central1.composer.cloud.google.com/jwt-public-key"
                },
                {
                  "name": "AIRFLOW__WEBSERVER__INVERTING_PROXY_BACKEND_ID",
                  "value": "2bddbf62981044cc9b7df0b2fb317261"
                },
                {
                  "name": "AIRFLOW__WEBSERVER__SECRET_KEY",
                  "value": "some-random-id"
                },
                {
                  "name": "AIRFLOW__WEBSERVER__UPDATE_FAB_PERMS",
                  "value": "True"
                },
                {
                  "name": "AIRFLOW__CORE__FERNET_KEY",
                  "valueFrom": {
                    "secretKeyRef": {
                      "key": "fernet_key",
                      "name": "airflow-secrets"
                    }
                  }
                },
                {
                  "name": "AIRFLOW__CORE__SQL_ALCHEMY_CONN",
                  "value": "postgresql+psycopg2://$(SQL_USER):$(SQL_PASSWORD)@airflow-sqlproxy-service.composer-system.svc.cluster.local:3306/$(SQL_DATABASE)"
                },
                {
                  "name": "CLOUDSDK_METRICS_ENVIRONMENT",
                  "value": "2.2.5+composer"
                },
                {
                  "name": "GCS_BUCKET",
                  "value": "us-central1-airflow-bed-152ea37f-bucket"
                },
                {
                  "name": "AIRFLOW_HOME",
                  "value": "/etc/airflow"
                },
                {
                  "name": "DAGS_FOLDER",
                  "value": "/home/airflow/gcs/dags"
                },
                {
                  "name": "GCSFUSE_EXTRACTED",
                  "value": "TRUE"
                },
                {
                  "name": "GCP_PROJECT",
                  "value": "digraph-2021"
                },
                {
                  "name": "COMPOSER_LOCATION",
                  "value": "us-central1"
                },
                {
                  "name": "COMPOSER_PYTHON_VERSION",
                  "value": "3"
                },
                {
                  "name": "COMPOSER_ENVIRONMENT",
                  "value": "airflow-bed"
                },
                {
                  "name": "GKE_CLUSTER_NAME",
                  "value": "us-central1-airflow-bed-152ea37f-gke"
                },
                {
                  "name": "CONTAINER_NAME",
                  "value": "airflow-webserver"
                }
              ],
              "image": "us-central1-docker.pkg.dev/digraph-2021/composer-images-us-central1-airflow-bed-152ea37f-gke/33e72803-04e7-4edc-ab9c-b7e19661b109",
              "imagePullPolicy": "IfNotPresent",
              "livenessProbe": {
                "exec": {
                  "command": [
                    "curl",
                    "--max-time",
                    "30",
                    "localhost:8080/_ah/health"
                  ]
                },
                "failureThreshold": 3,
                "periodSeconds": 10,
                "successThreshold": 1,
                "timeoutSeconds": 40
              },
              "name": "airflow-webserver",
              "resources": {
                "limits": {
                  "cpu": "700m",
                  "ephemeral-storage": "1843Mi",
                  "memory": "1433Mi"
                },
                "requests": {
                  "cpu": "700m",
                  "ephemeral-storage": "1843Mi",
                  "memory": "1433Mi"
                }
              },
              "securityContext": {
                "capabilities": {
                  "drop": [
                    "NET_RAW"
                  ]
                }
              },
              "startupProbe": {
                "exec": {
                  "command": [
                    "curl",
                    "--max-time",
                    "30",
                    "localhost:8080/_ah/health"
                  ]
                },
                "failureThreshold": 18,
                "periodSeconds": 10,
                "successThreshold": 1,
                "timeoutSeconds": 40
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File",
              "volumeMounts": [
                {
                  "mountPath": "/etc/airflow/airflow_cfg",
                  "name": "airflow-config"
                },
                {
                  "mountPath": "/home/airflow/gcs",
                  "name": "gcsdir"
                },
                {
                  "mountPath": "/home/airflow/container-comms",
                  "name": "container-comms"
                },
                {
                  "mountPath": "/home/airflow/gcsfuse",
                  "mountPropagation": "HostToContainer",
                  "name": "gcsfuse"
                }
              ]
            },
            {
              "args": [
                "/home/airflow/gcs"
              ],
              "env": [
                {
                  "name": "GCS_BUCKET",
                  "value": "us-central1-airflow-bed-152ea37f-bucket"
                },
                {
                  "name": "SQL_DATABASE",
                  "value": "composer-2-0-19-airflow-2-2-5-e4894b1f"
                },
                {
                  "name": "SQL_USER",
                  "value": "root"
                },
                {
                  "name": "SQL_PASSWORD",
                  "valueFrom": {
                    "secretKeyRef": {
                      "key": "sql_password",
                      "name": "airflow-secrets"
                    }
                  }
                },
                {
                  "name": "COMPOSER_GKE_NAME",
                  "value": "us-central1-airflow-bed-152ea37f-gke"
                },
                {
                  "name": "SKIP_SYNCING_DAGS",
                  "value": "TRUE"
                },
                {
                  "name": "AUTOGKE",
                  "value": "TRUE"
                },
                {
                  "name": "COMPOSER_GKE_LOCATION",
                  "value": "us-central1"
                }
              ],
              "image": "us-docker.pkg.dev/cloud-airflow-releaser/gcs-syncd/gcs-syncd:cloud_composer_service_2022-06-28-RC0",
              "imagePullPolicy": "IfNotPresent",
              "name": "gcs-syncd",
              "resources": {
                "limits": {
                  "cpu": "150m",
                  "ephemeral-storage": "102Mi",
                  "memory": "307Mi"
                },
                "requests": {
                  "cpu": "150m",
                  "ephemeral-storage": "102Mi",
                  "memory": "307Mi"
                }
              },
              "securityContext": {
                "capabilities": {
                  "drop": [
                    "NET_RAW"
                  ]
                }
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File",
              "volumeMounts": [
                {
                  "mountPath": "/home/airflow/gcs",
                  "name": "gcsdir"
                }
              ]
            },
            {
              "args": [
                "--backend=2bddbf62981044cc9b7df0b2fb317261",
                "--proxy=https://us-central1.composer.cloud.google.com/tun/m/4592f092208ecc84946b8f8f8016274df1b36a14/",
                "--shim-websockets=true",
                "--shim-path=websocket-shim",
                "--forward-user-id=true",
                "--health-check-path=/_ah/health",
                "--health-check-interval-seconds=5"
              ],
              "image": "us-docker.pkg.dev/cloud-airflow-releaser/composer-inverting-proxy/composer-inverting-proxy:cloud_composer_service_2022-06-28-RC0",
              "imagePullPolicy": "IfNotPresent",
              "name": "agent",
              "resources": {
                "limits": {
                  "cpu": "150m",
                  "ephemeral-storage": "102Mi",
                  "memory": "307Mi"
                },
                "requests": {
                  "cpu": "150m",
                  "ephemeral-storage": "102Mi",
                  "memory": "307Mi"
                }
              },
              "securityContext": {
                "capabilities": {
                  "drop": [
                    "NET_RAW"
                  ]
                }
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File"
            }
          ],
          "dnsPolicy": "ClusterFirst",
          "restartPolicy": "Always",
          "schedulerName": "default-scheduler",
          "securityContext": {
            "seccompProfile": {
              "type": "RuntimeDefault"
            }
          },
          "terminationGracePeriodSeconds": 30,
          "volumes": [
            {
              "configMap": {
                "defaultMode": 420,
                "name": "airflow-configmap"
              },
              "name": "airflow-config"
            },
            {
              "emptyDir": {},
              "name": "gcsdir"
            },
            {
              "hostPath": {
                "path": "/var/composer/gcs_mount_status",
                "type": ""
              },
              "name": "container-comms"
            },
            {
              "hostPath": {
                "path": "/var/composer/gcs_mount",
                "type": ""
              },
              "name": "gcsfuse"
            }
          ]
        }
      }
    },
    "status": {
      "availableReplicas": 1,
      "conditions": [
        {
          "lastTransitionTime": "2022-07-18T19:54:14Z",
          "lastUpdateTime": "2022-09-07T22:53:26Z",
          "message": "ReplicaSet \"airflow-webserver-5cd558768\" has successfully progressed.",
          "reason": "NewReplicaSetAvailable",
          "status": "True",
          "type": "Progressing"
        },
        {
          "lastTransitionTime": "2023-02-26T05:10:15Z",
          "lastUpdateTime": "2023-02-26T05:10:15Z",
          "message": "Deployment has minimum availability.",
          "reason": "MinimumReplicasAvailable",
          "status": "True",
          "type": "Available"
        }
      ],
      "observedGeneration": 4,
      "readyReplicas": 1,
      "replicas": 1,
      "updatedReplicas": 1
    }
  },
  {
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
      "annotations": {
        "deployment.kubernetes.io/revision": "2"
      },
      "creationTimestamp": "2022-02-15T05:55:21Z",
      "generation": 2,
      "labels": {
        "run": "airflow-monitoring"
      },
      "name": "airflow-monitoring",
      "namespace": "composer-system",
      "resourceVersion": "310754235",
      "uid": "919e0f15-eecb-4984-bf03-577b228e9c22"
    },
    "spec": {
      "progressDeadlineSeconds": 600,
      "replicas": 1,
      "revisionHistoryLimit": 10,
      "selector": {
        "matchLabels": {
          "run": "airflow-monitoring"
        }
      },
      "strategy": {
        "rollingUpdate": {
          "maxSurge": 1,
          "maxUnavailable": 1
        },
        "type": "RollingUpdate"
      },
      "template": {
        "metadata": {
          "creationTimestamp": null,
          "labels": {
            "composer-system-pod": "true",
            "run": "airflow-monitoring"
          }
        },
        "spec": {
          "containers": [
            {
              "args": [
                "/var/local/setup_monitoring.sh"
              ],
              "command": [
                "/bin/bash",
                "-ce"
              ],
              "env": [
                {
                  "name": "GCP_PROJECT",
                  "value": "digraph-2021"
                },
                {
                  "name": "COMPOSER_LOCATION",
                  "value": "us-central1"
                },
                {
                  "name": "COMPOSER_GKE_ZONE"
                },
                {
                  "name": "AUTOGKE",
                  "value": "TRUE"
                },
                {
                  "name": "COMPOSER_GKE_LOCATION",
                  "value": "us-central1"
                },
                {
                  "name": "GCS_BUCKET",
                  "value": "us-central1-airflow-bed-152ea37f-bucket"
                },
                {
                  "name": "COMPOSER_PYTHON_VERSION",
                  "value": "3"
                },
                {
                  "name": "REDIS_HOST",
                  "value": "airflow-redis-service.composer-system"
                },
                {
                  "name": "AIRFLOW_DATABASE_VERSION",
                  "value": "POSTGRES_13"
                },
                {
                  "name": "SQL_HOST",
                  "value": "airflow-sqlproxy-service.composer-system.svc.cluster.local"
                },
                {
                  "name": "SQL_DATABASE",
                  "value": "composer-2-0-19-airflow-2-2-5-e4894b1f"
                },
                {
                  "name": "COMPOSER_ENVIRONMENT",
                  "value": "airflow-bed"
                },
                {
                  "name": "IMAGE_VERSION",
                  "value": "composer-2-0-19-airflow-2-2-5"
                },
                {
                  "name": "GKE_CLUSTER_NAME",
                  "value": "us-central1-airflow-bed-152ea37f-gke"
                },
                {
                  "name": "CONTAINER_NAME",
                  "value": "airflow-monitoring"
                },
                {
                  "name": "CELERY_APP_NAME",
                  "value": "airflow.executors.celery_executor"
                },
                {
                  "name": "DEFAULT_QUEUE",
                  "value": "default"
                },
                {
                  "name": "STATSD_PORT",
                  "value": "8125"
                },
                {
                  "name": "REDIS_PORT",
                  "value": "6379"
                },
                {
                  "name": "SQL_USER",
                  "value": "root"
                },
                {
                  "name": "SQL_PASSWORD",
                  "valueFrom": {
                    "secretKeyRef": {
                      "key": "sql_password",
                      "name": "airflow-secrets-default"
                    }
                  }
                },
                {
                  "name": "COMPOSER_ENVIRONMENT_TYPE",
                  "value": "PUBLIC_IP"
                },
                {
                  "name": "TENANT_PROJECT",
                  "value": "a06a05266fdf96c46p-tp"
                },
                {
                  "name": "SQL_INSTANCE_NAME",
                  "value": "a06a05266fdf96c46p-tp:us-central1-airflow-bed-152ea37f-sql"
                },
                {
                  "name": "AUTOSCALING_ENABLED",
                  "value": "TRUE"
                },
                {
                  "name": "VERSIONED_NAMESPACE",
                  "value": "composer-2-0-19-airflow-2-2-5-e4894b1f"
                }
              ],
              "image": "us-docker.pkg.dev/cloud-airflow-releaser/airflow-monitoring/airflow-monitoring:cloud_composer_service_2022-06-28-RC0",
              "imagePullPolicy": "IfNotPresent",
              "livenessProbe": {
                "exec": {
                  "command": [
                    "/home/airflow/metric_prober.py"
                  ]
                },
                "failureThreshold": 1,
                "initialDelaySeconds": 300,
                "periodSeconds": 180,
                "successThreshold": 1,
                "timeoutSeconds": 30
              },
              "name": "airflow-monitoring",
              "resources": {
                "limits": {
                  "cpu": "300m",
                  "ephemeral-storage": "100Mi",
                  "memory": "600Mi"
                },
                "requests": {
                  "cpu": "200m",
                  "ephemeral-storage": "100Mi",
                  "memory": "600Mi"
                }
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File"
            },
            {
              "args": [
                "/home/airflow/start_task_monitor.sh"
              ],
              "command": [
                "/bin/bash",
                "-ce"
              ],
              "env": [
                {
                  "name": "COMPOSER_GKE_ZONE"
                },
                {
                  "name": "COMPOSER_GKE_NAME",
                  "value": "us-central1-airflow-bed-152ea37f-gke"
                },
                {
                  "name": "AUTOGKE",
                  "value": "TRUE"
                },
                {
                  "name": "COMPOSER_GKE_LOCATION",
                  "value": "us-central1"
                },
                {
                  "name": "SQL_HOST",
                  "value": "airflow-sqlproxy-service.composer-system.svc.cluster.local"
                },
                {
                  "name": "SQL_USER",
                  "value": "root"
                },
                {
                  "name": "SQL_PASSWORD",
                  "valueFrom": {
                    "secretKeyRef": {
                      "key": "sql_password",
                      "name": "airflow-secrets-default"
                    }
                  }
                },
                {
                  "name": "SQL_DATABASE",
                  "value": "composer-2-0-19-airflow-2-2-5-e4894b1f"
                },
                {
                  "name": "AIRFLOW_DATABASE_VERSION",
                  "value": "POSTGRES_13"
                }
              ],
              "image": "us-docker.pkg.dev/cloud-airflow-releaser/task-monitor/task-monitor:cloud_composer_service_2022-06-28-RC0",
              "imagePullPolicy": "IfNotPresent",
              "name": "task-monitor",
              "resources": {
                "limits": {
                  "cpu": "50m",
                  "ephemeral-storage": "100Mi",
                  "memory": "100Mi"
                },
                "requests": {
                  "cpu": "50m",
                  "ephemeral-storage": "100Mi",
                  "memory": "100Mi"
                }
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File"
            }
          ],
          "dnsPolicy": "ClusterFirst",
          "restartPolicy": "Always",
          "schedulerName": "default-scheduler",
          "securityContext": {},
          "terminationGracePeriodSeconds": 30
        }
      }
    },
    "status": {
      "availableReplicas": 1,
      "conditions": [
        {
          "lastTransitionTime": "2022-02-15T05:55:22Z",
          "lastUpdateTime": "2022-02-15T05:55:22Z",
          "message": "Deployment has minimum availability.",
          "reason": "MinimumReplicasAvailable",
          "status": "True",
          "type": "Available"
        },
        {
          "lastTransitionTime": "2022-02-15T05:55:21Z",
          "lastUpdateTime": "2022-07-18T19:55:20Z",
          "message": "ReplicaSet \"airflow-monitoring-6f87f44ddf\" has successfully progressed.",
          "reason": "NewReplicaSetAvailable",
          "status": "True",
          "type": "Progressing"
        }
      ],
      "observedGeneration": 2,
      "readyReplicas": 1,
      "replicas": 1,
      "updatedReplicas": 1
    }
  },
  {
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
      "annotations": {
        "deployment.kubernetes.io/revision": "3"
      },
      "creationTimestamp": "2022-02-15T05:55:22Z",
      "generation": 16,
      "labels": {
        "run": "airflow-sqlproxy"
      },
      "name": "airflow-sqlproxy",
      "namespace": "composer-system",
      "resourceVersion": "284920807",
      "uid": "f305126a-ffe9-4092-b2d4-47275fe0c771"
    },
    "spec": {
      "minReadySeconds": 5,
      "progressDeadlineSeconds": 600,
      "replicas": 1,
      "revisionHistoryLimit": 10,
      "selector": {
        "matchLabels": {
          "run": "airflow-sqlproxy"
        }
      },
      "strategy": {
        "rollingUpdate": {
          "maxSurge": "50%",
          "maxUnavailable": 0
        },
        "type": "RollingUpdate"
      },
      "template": {
        "metadata": {
          "creationTimestamp": null,
          "labels": {
            "composer-system-pod": "true",
            "run": "airflow-sqlproxy"
          }
        },
        "spec": {
          "affinity": {
            "podAntiAffinity": {
              "preferredDuringSchedulingIgnoredDuringExecution": [
                {
                  "podAffinityTerm": {
                    "labelSelector": {
                      "matchLabels": {
                        "run": "airflow-scheduler"
                      }
                    },
                    "namespaces": [
                      "composer-2-0-19-airflow-2-2-5-e4894b1f"
                    ],
                    "topologyKey": "kubernetes.io/hostname"
                  },
                  "weight": 1
                }
              ]
            }
          },
          "containers": [
            {
              "env": [
                {
                  "name": "AIRFLOW_DATABASE_VERSION",
                  "value": "POSTGRES_13"
                },
                {
                  "name": "SQL_PROXY_INSTANCES",
                  "value": "a06a05266fdf96c46p-tp:us-central1:us-central1-airflow-bed-152ea37f-sql=tcp:0.0.0.0:3306"
                },
                {
                  "name": "SQL_PROXY_TERM_TIMEOUT",
                  "value": "585s"
                },
                {
                  "name": "SQL_PASSWORD",
                  "valueFrom": {
                    "secretKeyRef": {
                      "key": "sql_password",
                      "name": "airflow-secrets-default"
                    }
                  }
                },
                {
                  "name": "HEALTHCHECK_PORT",
                  "value": "3307"
                }
              ],
              "image": "us-docker.pkg.dev/cloud-airflow-releaser/composer-cloudsql-proxy/composer-cloudsql-proxy:cloud_composer_service_2022-06-28-RC0",
              "imagePullPolicy": "IfNotPresent",
              "livenessProbe": {
                "exec": {
                  "command": [
                    "/var/local/liveness_probe.sh"
                  ]
                },
                "failureThreshold": 3,
                "initialDelaySeconds": 540,
                "periodSeconds": 30,
                "successThreshold": 1,
                "timeoutSeconds": 30
              },
              "name": "airflow-sqlproxy",
              "ports": [
                {
                  "containerPort": 3306,
                  "protocol": "TCP"
                }
              ],
              "readinessProbe": {
                "exec": {
                  "command": [
                    "/var/local/liveness_probe.sh"
                  ]
                },
                "failureThreshold": 54,
                "initialDelaySeconds": 10,
                "periodSeconds": 10,
                "successThreshold": 1,
                "timeoutSeconds": 30
              },
              "resources": {
                "limits": {
                  "cpu": "200m",
                  "ephemeral-storage": "100Mi",
                  "memory": "200Mi"
                },
                "requests": {
                  "cpu": "200m",
                  "ephemeral-storage": "100Mi",
                  "memory": "200Mi"
                }
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File"
            }
          ],
          "dnsPolicy": "ClusterFirst",
          "priorityClassName": "highest-priority",
          "restartPolicy": "Always",
          "schedulerName": "default-scheduler",
          "securityContext": {},
          "terminationGracePeriodSeconds": 600
        }
      }
    },
    "status": {
      "availableReplicas": 1,
      "conditions": [
        {
          "lastTransitionTime": "2022-02-15T05:55:22Z",
          "lastUpdateTime": "2022-07-18T20:01:26Z",
          "message": "ReplicaSet \"airflow-sqlproxy-794c8bb57b\" has successfully progressed.",
          "reason": "NewReplicaSetAvailable",
          "status": "True",
          "type": "Progressing"
        },
        {
          "lastTransitionTime": "2023-02-26T05:13:20Z",
          "lastUpdateTime": "2023-02-26T05:13:20Z",
          "message": "Deployment has minimum availability.",
          "reason": "MinimumReplicasAvailable",
          "status": "True",
          "type": "Available"
        }
      ],
      "observedGeneration": 16,
      "readyReplicas": 1,
      "replicas": 1,
      "updatedReplicas": 1
    }
  },
  {
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
      "annotations": {
        "deployment.kubernetes.io/revision": "1"
      },
      "creationTimestamp": "2022-02-15T05:55:22Z",
      "generation": 1,
      "labels": {
        "control-plane": "worker-set-controller"
      },
      "name": "airflow-worker-set-controller",
      "namespace": "composer-system",
      "resourceVersion": "284150870",
      "uid": "84aa33d2-73a0-41e1-aaed-2bf57c54082f"
    },
    "spec": {
      "progressDeadlineSeconds": 600,
      "replicas": 1,
      "revisionHistoryLimit": 10,
      "selector": {
        "matchLabels": {
          "control-plane": "worker-set-controller"
        }
      },
      "strategy": {
        "type": "Recreate"
      },
      "template": {
        "metadata": {
          "annotations": {
            "cluster-autoscaler.kubernetes.io/safe-to-evict": "true"
          },
          "creationTimestamp": null,
          "labels": {
            "control-plane": "worker-set-controller"
          }
        },
        "spec": {
          "containers": [
            {
              "args": [
                "--metrics-addr=127.0.0.1:8080"
              ],
              "command": [
                "/manager"
              ],
              "image": "us-docker.pkg.dev/cloud-airflow-releaser/airflow-worker-set-controller/airflow-worker-set-controller:cloud_composer_service_2022-02-01-RC1",
              "imagePullPolicy": "IfNotPresent",
              "name": "manager",
              "resources": {},
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File"
            },
            {
              "args": [
                "--secure-listen-address=0.0.0.0:8443",
                "--upstream=http://127.0.0.1:8080/",
                "--logtostderr=true",
                "--v=10"
              ],
              "image": "gcr.io/kubebuilder/kube-rbac-proxy:v0.5.0",
              "imagePullPolicy": "IfNotPresent",
              "name": "kube-rbac-proxy",
              "ports": [
                {
                  "containerPort": 8443,
                  "name": "https",
                  "protocol": "TCP"
                }
              ],
              "resources": {
                "limits": {
                  "cpu": "50m",
                  "ephemeral-storage": "100Mi",
                  "memory": "50Mi"
                },
                "requests": {
                  "cpu": "50m",
                  "ephemeral-storage": "100Mi",
                  "memory": "50Mi"
                }
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File"
            }
          ],
          "dnsPolicy": "ClusterFirst",
          "restartPolicy": "Always",
          "schedulerName": "default-scheduler",
          "securityContext": {},
          "terminationGracePeriodSeconds": 10
        }
      }
    },
    "status": {
      "availableReplicas": 1,
      "conditions": [
        {
          "lastTransitionTime": "2022-02-15T05:55:23Z",
          "lastUpdateTime": "2022-02-15T05:55:29Z",
          "message": "ReplicaSet \"airflow-worker-set-controller-f8b6c4c9b\" has successfully progressed.",
          "reason": "NewReplicaSetAvailable",
          "status": "True",
          "type": "Progressing"
        },
        {
          "lastTransitionTime": "2023-02-25T05:14:10Z",
          "lastUpdateTime": "2023-02-25T05:14:10Z",
          "message": "Deployment has minimum availability.",
          "reason": "MinimumReplicasAvailable",
          "status": "True",
          "type": "Available"
        }
      ],
      "observedGeneration": 1,
      "readyReplicas": 1,
      "replicas": 1,
      "updatedReplicas": 1
    }
  },
  {
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
      "annotations": {
        "deployment.kubernetes.io/revision": "3"
      },
      "creationTimestamp": "2022-09-05T13:30:32Z",
      "generation": 3,
      "labels": {
        "composer-infra": "critical",
        "run": "autohealing"
      },
      "name": "composer-autohealing",
      "namespace": "composer-system",
      "resourceVersion": "284150802",
      "uid": "47ee47e1-bf20-469c-82ac-14ffe082982d"
    },
    "spec": {
      "progressDeadlineSeconds": 600,
      "replicas": 1,
      "revisionHistoryLimit": 10,
      "selector": {
        "matchLabels": {
          "run": "autohealing"
        }
      },
      "strategy": {
        "rollingUpdate": {
          "maxSurge": "25%",
          "maxUnavailable": "25%"
        },
        "type": "RollingUpdate"
      },
      "template": {
        "metadata": {
          "creationTimestamp": null,
          "labels": {
            "composer-infra": "critical",
            "run": "autohealing"
          }
        },
        "spec": {
          "containers": [
            {
              "env": [
                {
                  "name": "COMPOSER_V2",
                  "value": "TRUE"
                },
                {
                  "name": "AUTOHEALER_NO_DRY_RUN_ACTIONS",
                  "value": "DELETE_ANETD_POD"
                },
                {
                  "name": "PROJECT_ID",
                  "value": "digraph-2021"
                },
                {
                  "name": "LOCATION",
                  "value": "us-central1"
                },
                {
                  "name": "ENVIRONMENT_NAME",
                  "value": "airflow-bed"
                },
                {
                  "name": "ENVIRONMENT_UUID",
                  "value": "2bddbf62-9810-44cc-9b7d-f0b2fb317261"
                },
                {
                  "name": "HEARTBEAT_FILEPATH",
                  "value": "/tmp/heartbeat"
                }
              ],
              "image": "us-docker.pkg.dev/cloud-airflow-releaser/composer-autohealing/composer-autohealing:dnsfix_rc4",
              "imagePullPolicy": "IfNotPresent",
              "livenessProbe": {
                "exec": {
                  "command": [
                    "/opt/probe"
                  ]
                },
                "failureThreshold": 3,
                "initialDelaySeconds": 60,
                "periodSeconds": 180,
                "successThreshold": 1,
                "timeoutSeconds": 30
              },
              "name": "composer-autohealing",
              "resources": {},
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File"
            }
          ],
          "dnsPolicy": "ClusterFirst",
          "restartPolicy": "Always",
          "schedulerName": "default-scheduler",
          "securityContext": {},
          "serviceAccount": "privileged",
          "serviceAccountName": "privileged",
          "terminationGracePeriodSeconds": 90
        }
      }
    },
    "status": {
      "availableReplicas": 1,
      "conditions": [
        {
          "lastTransitionTime": "2022-09-05T13:30:32Z",
          "lastUpdateTime": "2023-02-12T12:40:23Z",
          "message": "ReplicaSet \"composer-autohealing-68c9f5db8b\" has successfully progressed.",
          "reason": "NewReplicaSetAvailable",
          "status": "True",
          "type": "Progressing"
        },
        {
          "lastTransitionTime": "2023-02-25T05:14:05Z",
          "lastUpdateTime": "2023-02-25T05:14:05Z",
          "message": "Deployment has minimum availability.",
          "reason": "MinimumReplicasAvailable",
          "status": "True",
          "type": "Available"
        }
      ],
      "observedGeneration": 3,
      "readyReplicas": 1,
      "replicas": 1,
      "updatedReplicas": 1
    }
  },
  {
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
      "annotations": {
        "deployment.kubernetes.io/revision": "2",
        "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"apps/v1\",\"kind\":\"Deployment\",\"metadata\":{\"annotations\":{},\"labels\":{\"k8s-app\":\"custom-metrics-stackdriver-adapter\",\"run\":\"custom-metrics-stackdriver-adapter\"},\"name\":\"custom-metrics-stackdriver-adapter\",\"namespace\":\"composer-system\"},\"spec\":{\"replicas\":1,\"selector\":{\"matchLabels\":{\"k8s-app\":\"custom-metrics-stackdriver-adapter\",\"run\":\"custom-metrics-stackdriver-adapter\"}},\"template\":{\"metadata\":{\"annotations\":{\"cluster-autoscaler.kubernetes.io/safe-to-evict\":\"true\"},\"labels\":{\"k8s-app\":\"custom-metrics-stackdriver-adapter\",\"kubernetes.io/cluster-service\":\"true\",\"run\":\"custom-metrics-stackdriver-adapter\"}},\"spec\":{\"containers\":[{\"command\":[\"/adapter\",\"--use-new-resource-model=true\"],\"image\":\"gcr.io/gke-release/custom-metrics-stackdriver-adapter:v0.12.0-gke.0\",\"imagePullPolicy\":\"Always\",\"name\":\"pod-custom-metrics-stackdriver-adapter\",\"resources\":{\"limits\":{\"cpu\":\"100m\",\"ephemeral-storage\":\"100Mi\",\"memory\":\"200Mi\"},\"requests\":{\"cpu\":\"50m\",\"ephemeral-storage\":\"100Mi\",\"memory\":\"200Mi\"}}}],\"serviceAccountName\":\"custom-metrics-stackdriver-adapter\"}}}}\n"
      },
      "creationTimestamp": "2022-02-15T05:55:15Z",
      "generation": 2,
      "labels": {
        "k8s-app": "custom-metrics-stackdriver-adapter",
        "run": "custom-metrics-stackdriver-adapter"
      },
      "name": "custom-metrics-stackdriver-adapter",
      "namespace": "composer-system",
      "resourceVersion": "284886570",
      "uid": "4a825101-f166-4a2b-ada8-07cb1ad6e82e"
    },
    "spec": {
      "progressDeadlineSeconds": 600,
      "replicas": 1,
      "revisionHistoryLimit": 10,
      "selector": {
        "matchLabels": {
          "k8s-app": "custom-metrics-stackdriver-adapter",
          "run": "custom-metrics-stackdriver-adapter"
        }
      },
      "strategy": {
        "rollingUpdate": {
          "maxSurge": "25%",
          "maxUnavailable": "25%"
        },
        "type": "RollingUpdate"
      },
      "template": {
        "metadata": {
          "annotations": {
            "cluster-autoscaler.kubernetes.io/safe-to-evict": "true"
          },
          "creationTimestamp": null,
          "labels": {
            "k8s-app": "custom-metrics-stackdriver-adapter",
            "kubernetes.io/cluster-service": "true",
            "run": "custom-metrics-stackdriver-adapter"
          }
        },
        "spec": {
          "containers": [
            {
              "command": [
                "/adapter",
                "--use-new-resource-model=true"
              ],
              "image": "gcr.io/gke-release/custom-metrics-stackdriver-adapter:v0.12.2-gke.0",
              "imagePullPolicy": "Always",
              "name": "pod-custom-metrics-stackdriver-adapter",
              "resources": {
                "limits": {
                  "cpu": "250m",
                  "ephemeral-storage": "102Mi",
                  "memory": "409Mi"
                },
                "requests": {
                  "cpu": "250m",
                  "ephemeral-storage": "102Mi",
                  "memory": "409Mi"
                }
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File"
            }
          ],
          "dnsPolicy": "ClusterFirst",
          "restartPolicy": "Always",
          "schedulerName": "default-scheduler",
          "securityContext": {},
          "serviceAccount": "custom-metrics-stackdriver-adapter",
          "serviceAccountName": "custom-metrics-stackdriver-adapter",
          "terminationGracePeriodSeconds": 30
        }
      }
    },
    "status": {
      "availableReplicas": 1,
      "conditions": [
        {
          "lastTransitionTime": "2022-02-15T05:55:15Z",
          "lastUpdateTime": "2022-06-13T10:26:52Z",
          "message": "ReplicaSet \"custom-metrics-stackdriver-adapter-7b5f8c6fc\" has successfully progressed.",
          "reason": "NewReplicaSetAvailable",
          "status": "True",
          "type": "Progressing"
        },
        {
          "lastTransitionTime": "2023-02-26T04:10:41Z",
          "lastUpdateTime": "2023-02-26T04:10:41Z",
          "message": "Deployment has minimum availability.",
          "reason": "MinimumReplicasAvailable",
          "status": "True",
          "type": "Available"
        }
      ],
      "observedGeneration": 2,
      "readyReplicas": 1,
      "replicas": 1,
      "updatedReplicas": 1
    }
  },
  {
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
      "annotations": {
        "deployment.kubernetes.io/revision": "2"
      },
      "creationTimestamp": "2022-02-15T05:46:18Z",
      "generation": 3,
      "labels": {
        "addonmanager.kubernetes.io/mode": "Reconcile",
        "k8s-app": "event-exporter",
        "kubernetes.io/cluster-service": "true",
        "version": "v0.3.9"
      },
      "name": "event-exporter-gke",
      "namespace": "kube-system",
      "resourceVersion": "305781635",
      "uid": "60bd485e-8676-47db-a5c4-134aa9b4b20c"
    },
    "spec": {
      "progressDeadlineSeconds": 600,
      "replicas": 1,
      "revisionHistoryLimit": 10,
      "selector": {
        "matchLabels": {
          "k8s-app": "event-exporter"
        }
      },
      "strategy": {
        "rollingUpdate": {
          "maxSurge": "25%",
          "maxUnavailable": "25%"
        },
        "type": "RollingUpdate"
      },
      "template": {
        "metadata": {
          "annotations": {
            "components.gke.io/component-name": "event-exporter",
            "components.gke.io/component-version": "1.1.0"
          },
          "creationTimestamp": null,
          "labels": {
            "k8s-app": "event-exporter",
            "version": "v0.3.9"
          }
        },
        "spec": {
          "containers": [
            {
              "command": [
                "/event-exporter",
                "-sink-opts=-stackdriver-resource-model=new -endpoint=https://logging.googleapis.com",
                "-prometheus-endpoint=:8080"
              ],
              "image": "gke.gcr.io/event-exporter:v0.3.9-gke.0",
              "imagePullPolicy": "IfNotPresent",
              "name": "event-exporter",
              "resources": {},
              "securityContext": {
                "allowPrivilegeEscalation": false,
                "capabilities": {
                  "drop": [
                    "all"
                  ]
                }
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File"
            },
            {
              "command": [
                "/monitor",
                "--stackdriver-prefix=container.googleapis.com/internal/addons",
                "--api-override=https://monitoring.googleapis.com/",
                "--source=event_exporter:http://localhost:8080?whitelisted=stackdriver_sink_received_entry_count,stackdriver_sink_request_count,stackdriver_sink_successfully_sent_entry_count",
                "--pod-id=$(POD_NAME)",
                "--namespace-id=$(POD_NAMESPACE)",
                "--node-name=$(NODE_NAME)"
              ],
              "env": [
                {
                  "name": "POD_NAME",
                  "valueFrom": {
                    "fieldRef": {
                      "apiVersion": "v1",
                      "fieldPath": "metadata.name"
                    }
                  }
                },
                {
                  "name": "POD_NAMESPACE",
                  "valueFrom": {
                    "fieldRef": {
                      "apiVersion": "v1",
                      "fieldPath": "metadata.namespace"
                    }
                  }
                },
                {
                  "name": "NODE_NAME",
                  "valueFrom": {
                    "fieldRef": {
                      "apiVersion": "v1",
                      "fieldPath": "spec.nodeName"
                    }
                  }
                }
              ],
              "image": "gke.gcr.io/prometheus-to-sd:v0.10.0-gke.0",
              "imagePullPolicy": "IfNotPresent",
              "name": "prometheus-to-sd-exporter",
              "resources": {},
              "securityContext": {
                "allowPrivilegeEscalation": false,
                "capabilities": {
                  "drop": [
                    "all"
                  ]
                }
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File"
            }
          ],
          "dnsPolicy": "ClusterFirst",
          "nodeSelector": {
            "kubernetes.io/os": "linux"
          },
          "restartPolicy": "Always",
          "schedulerName": "default-scheduler",
          "securityContext": {
            "runAsGroup": 1000,
            "runAsUser": 1000
          },
          "serviceAccount": "event-exporter-sa",
          "serviceAccountName": "event-exporter-sa",
          "terminationGracePeriodSeconds": 120,
          "tolerations": [
            {
              "key": "components.gke.io/gke-managed-components",
              "operator": "Exists"
            }
          ],
          "volumes": [
            {
              "hostPath": {
                "path": "/etc/ssl/certs",
                "type": "Directory"
              },
              "name": "ssl-certs"
            }
          ]
        }
      }
    },
    "status": {
      "availableReplicas": 1,
      "conditions": [
        {
          "lastTransitionTime": "2022-02-15T05:46:18Z",
          "lastUpdateTime": "2022-11-15T02:20:41Z",
          "message": "ReplicaSet \"event-exporter-gke-f66d9f855\" has successfully progressed.",
          "reason": "NewReplicaSetAvailable",
          "status": "True",
          "type": "Progressing"
        },
        {
          "lastTransitionTime": "2023-02-25T05:14:15Z",
          "lastUpdateTime": "2023-02-25T05:14:15Z",
          "message": "Deployment has minimum availability.",
          "reason": "MinimumReplicasAvailable",
          "status": "True",
          "type": "Available"
        }
      ],
      "observedGeneration": 3,
      "readyReplicas": 1,
      "replicas": 1,
      "updatedReplicas": 1
    }
  },
  {
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
      "annotations": {
        "components.gke.io/layer": "addon",
        "credential-normal-mode": "true",
        "deployment.kubernetes.io/revision": "5"
      },
      "creationTimestamp": "2022-02-15T05:46:34Z",
      "generation": 82,
      "labels": {
        "addonmanager.kubernetes.io/mode": "Reconcile",
        "k8s-app": "konnectivity-agent"
      },
      "name": "konnectivity-agent",
      "namespace": "kube-system",
      "resourceVersion": "305781720",
      "uid": "5859a5ec-9ef1-4e4c-b2d9-db8de54b9da6"
    },
    "spec": {
      "progressDeadlineSeconds": 600,
      "replicas": 3,
      "revisionHistoryLimit": 10,
      "selector": {
        "matchLabels": {
          "k8s-app": "konnectivity-agent"
        }
      },
      "strategy": {
        "rollingUpdate": {
          "maxSurge": "25%",
          "maxUnavailable": "25%"
        },
        "type": "RollingUpdate"
      },
      "template": {
        "metadata": {
          "annotations": {
            "cluster-autoscaler.kubernetes.io/safe-to-evict": "true",
            "components.gke.io/component-name": "konnectivitynetworkproxy-combined",
            "components.gke.io/component-version": "1.4.10"
          },
          "creationTimestamp": null,
          "labels": {
            "k8s-app": "konnectivity-agent"
          }
        },
        "spec": {
          "containers": [
            {
              "args": [
                "--logtostderr=true",
                "--ca-cert=/var/run/secrets/kubernetes.io/serviceaccount/ca.crt",
                "--proxy-server-host=35.225.148.2",
                "--proxy-server-port=8132",
                "--health-server-port=8093",
                "--admin-server-port=8094",
                "--sync-interval=5s",
                "--sync-interval-cap=30s",
                "--sync-forever=true",
                "--probe-interval=5s",
                "--keepalive-time=60s",
                "--service-account-token-path=/var/run/secrets/tokens/konnectivity-agent-token",
                "--v=3"
              ],
              "command": [
                "/proxy-agent"
              ],
              "env": [
                {
                  "name": "POD_NAME",
                  "valueFrom": {
                    "fieldRef": {
                      "apiVersion": "v1",
                      "fieldPath": "metadata.name"
                    }
                  }
                },
                {
                  "name": "POD_NAMESPACE",
                  "valueFrom": {
                    "fieldRef": {
                      "apiVersion": "v1",
                      "fieldPath": "metadata.namespace"
                    }
                  }
                }
              ],
              "image": "gke.gcr.io/proxy-agent:v0.0.31-gke.0",
              "imagePullPolicy": "IfNotPresent",
              "livenessProbe": {
                "failureThreshold": 3,
                "httpGet": {
                  "path": "/healthz",
                  "port": 8093,
                  "scheme": "HTTP"
                },
                "initialDelaySeconds": 15,
                "periodSeconds": 10,
                "successThreshold": 1,
                "timeoutSeconds": 15
              },
              "name": "konnectivity-agent",
              "ports": [
                {
                  "containerPort": 8093,
                  "name": "metrics",
                  "protocol": "TCP"
                }
              ],
              "resources": {
                "limits": {
                  "memory": "125Mi"
                },
                "requests": {
                  "cpu": "10m",
                  "memory": "30Mi"
                }
              },
              "securityContext": {
                "allowPrivilegeEscalation": false,
                "capabilities": {
                  "drop": [
                    "all"
                  ]
                }
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File",
              "volumeMounts": [
                {
                  "mountPath": "/var/run/secrets/tokens",
                  "name": "konnectivity-agent-token"
                }
              ]
            }
          ],
          "dnsPolicy": "ClusterFirst",
          "nodeSelector": {
            "beta.kubernetes.io/os": "linux"
          },
          "priorityClassName": "system-cluster-critical",
          "restartPolicy": "Always",
          "schedulerName": "default-scheduler",
          "securityContext": {
            "fsGroup": 1000,
            "runAsGroup": 1000,
            "runAsUser": 1000
          },
          "serviceAccount": "konnectivity-agent",
          "serviceAccountName": "konnectivity-agent",
          "terminationGracePeriodSeconds": 30,
          "tolerations": [
            {
              "key": "CriticalAddonsOnly",
              "operator": "Exists"
            },
            {
              "effect": "NoSchedule",
              "key": "sandbox.gke.io/runtime",
              "operator": "Equal",
              "value": "gvisor"
            },
            {
              "key": "components.gke.io/gke-managed-components",
              "operator": "Exists"
            }
          ],
          "topologySpreadConstraints": [
            {
              "labelSelector": {
                "matchLabels": {
                  "k8s-app": "konnectivity-agent"
                }
              },
              "maxSkew": 1,
              "topologyKey": "topology.kubernetes.io/zone",
              "whenUnsatisfiable": "ScheduleAnyway"
            },
            {
              "labelSelector": {
                "matchLabels": {
                  "k8s-app": "konnectivity-agent"
                }
              },
              "maxSkew": 1,
              "topologyKey": "kubernetes.io/hostname",
              "whenUnsatisfiable": "ScheduleAnyway"
            }
          ],
          "volumes": [
            {
              "name": "konnectivity-agent-token",
              "projected": {
                "defaultMode": 420,
                "sources": [
                  {
                    "serviceAccountToken": {
                      "audience": "system:konnectivity-server",
                      "expirationSeconds": 3600,
                      "path": "konnectivity-agent-token"
                    }
                  }
                ]
              }
            }
          ]
        }
      }
    },
    "status": {
      "availableReplicas": 3,
      "conditions": [
        {
          "lastTransitionTime": "2022-02-15T05:46:34Z",
          "lastUpdateTime": "2022-09-28T09:26:49Z",
          "message": "ReplicaSet \"konnectivity-agent-68b565957b\" has successfully progressed.",
          "reason": "NewReplicaSetAvailable",
          "status": "True",
          "type": "Progressing"
        },
        {
          "lastTransitionTime": "2023-02-26T05:12:51Z",
          "lastUpdateTime": "2023-02-26T05:12:51Z",
          "message": "Deployment has minimum availability.",
          "reason": "MinimumReplicasAvailable",
          "status": "True",
          "type": "Available"
        }
      ],
      "observedGeneration": 82,
      "readyReplicas": 3,
      "replicas": 3,
      "updatedReplicas": 3
    }
  },
  {
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
      "annotations": {
        "components.gke.io/layer": "addon",
        "deployment.kubernetes.io/revision": "5"
      },
      "creationTimestamp": "2022-02-15T05:46:19Z",
      "generation": 6,
      "labels": {
        "addonmanager.kubernetes.io/mode": "Reconcile",
        "k8s-app": "konnectivity-agent-autoscaler",
        "kubernetes.io/cluster-service": "true"
      },
      "name": "konnectivity-agent-autoscaler",
      "namespace": "kube-system",
      "resourceVersion": "284150214",
      "uid": "d7e6aea3-9991-423b-97fa-04fd187bc2e6"
    },
    "spec": {
      "progressDeadlineSeconds": 600,
      "replicas": 1,
      "revisionHistoryLimit": 10,
      "selector": {
        "matchLabels": {
          "k8s-app": "konnectivity-agent-autoscaler"
        }
      },
      "strategy": {
        "rollingUpdate": {
          "maxSurge": "25%",
          "maxUnavailable": "25%"
        },
        "type": "RollingUpdate"
      },
      "template": {
        "metadata": {
          "annotations": {
            "cluster-autoscaler.kubernetes.io/safe-to-evict": "true",
            "components.gke.io/component-name": "konnectivitynetworkproxy-combined",
            "components.gke.io/component-version": "1.4.10"
          },
          "creationTimestamp": null,
          "labels": {
            "k8s-app": "konnectivity-agent-autoscaler"
          }
        },
        "spec": {
          "containers": [
            {
              "command": [
                "/cluster-proportional-autoscaler",
                "--namespace=kube-system",
                "--configmap=konnectivity-agent-autoscaler-config",
                "--target=deployment/konnectivity-agent",
                "--logtostderr=true",
                "--v=2"
              ],
              "image": "gke.gcr.io/cluster-proportional-autoscaler:1.8.4-gke.1",
              "imagePullPolicy": "IfNotPresent",
              "name": "autoscaler",
              "resources": {
                "requests": {
                  "cpu": "10m",
                  "memory": "10M"
                }
              },
              "securityContext": {
                "allowPrivilegeEscalation": false,
                "capabilities": {
                  "drop": [
                    "all"
                  ]
                }
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File"
            }
          ],
          "dnsPolicy": "ClusterFirst",
          "nodeSelector": {
            "beta.kubernetes.io/os": "linux"
          },
          "priorityClassName": "system-cluster-critical",
          "restartPolicy": "Always",
          "schedulerName": "default-scheduler",
          "securityContext": {
            "runAsGroup": 1000,
            "runAsUser": 1000
          },
          "serviceAccount": "konnectivity-agent-cpha",
          "serviceAccountName": "konnectivity-agent-cpha",
          "terminationGracePeriodSeconds": 30,
          "tolerations": [
            {
              "key": "CriticalAddonsOnly",
              "operator": "Exists"
            },
            {
              "key": "components.gke.io/gke-managed-components",
              "operator": "Exists"
            }
          ]
        }
      }
    },
    "status": {
      "availableReplicas": 1,
      "conditions": [
        {
          "lastTransitionTime": "2022-02-15T05:46:19Z",
          "lastUpdateTime": "2022-09-28T09:26:47Z",
          "message": "ReplicaSet \"konnectivity-agent-autoscaler-6dfb4f9cfb\" has successfully progressed.",
          "reason": "NewReplicaSetAvailable",
          "status": "True",
          "type": "Progressing"
        },
        {
          "lastTransitionTime": "2023-02-25T05:13:04Z",
          "lastUpdateTime": "2023-02-25T05:13:04Z",
          "message": "Deployment has minimum availability.",
          "reason": "MinimumReplicasAvailable",
          "status": "True",
          "type": "Available"
        }
      ],
      "observedGeneration": 6,
      "readyReplicas": 1,
      "replicas": 1,
      "updatedReplicas": 1
    }
  },
  {
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
      "annotations": {
        "deployment.kubernetes.io/revision": "6"
      },
      "creationTimestamp": "2022-02-15T05:46:17Z",
      "generation": 9,
      "labels": {
        "addonmanager.kubernetes.io/mode": "Reconcile",
        "k8s-app": "kube-dns",
        "kubernetes.io/cluster-service": "true"
      },
      "name": "kube-dns",
      "namespace": "kube-system",
      "resourceVersion": "305781578",
      "uid": "0db97a9d-06a8-4171-a75c-99c20663c8de"
    },
    "spec": {
      "progressDeadlineSeconds": 600,
      "replicas": 2,
      "revisionHistoryLimit": 10,
      "selector": {
        "matchLabels": {
          "k8s-app": "kube-dns"
        }
      },
      "strategy": {
        "rollingUpdate": {
          "maxSurge": "10%",
          "maxUnavailable": 0
        },
        "type": "RollingUpdate"
      },
      "template": {
        "metadata": {
          "annotations": {
            "components.gke.io/component-name": "kubedns",
            "prometheus.io/port": "10054",
            "prometheus.io/scrape": "true",
            "scheduler.alpha.kubernetes.io/critical-pod": "",
            "seccomp.security.alpha.kubernetes.io/pod": "runtime/default"
          },
          "creationTimestamp": null,
          "labels": {
            "k8s-app": "kube-dns"
          }
        },
        "spec": {
          "affinity": {
            "podAntiAffinity": {
              "preferredDuringSchedulingIgnoredDuringExecution": [
                {
                  "podAffinityTerm": {
                    "labelSelector": {
                      "matchExpressions": [
                        {
                          "key": "k8s-app",
                          "operator": "In",
                          "values": [
                            "kube-dns"
                          ]
                        }
                      ]
                    },
                    "topologyKey": "kubernetes.io/hostname"
                  },
                  "weight": 100
                }
              ]
            }
          },
          "containers": [
            {
              "args": [
                "--domain=cluster.local.",
                "--dns-port=10053",
                "--config-dir=/kube-dns-config",
                "--v=2"
              ],
              "env": [
                {
                  "name": "PROMETHEUS_PORT",
                  "value": "10055"
                }
              ],
              "image": "gke.gcr.io/k8s-dns-kube-dns:1.22.12-gke.0",
              "imagePullPolicy": "IfNotPresent",
              "livenessProbe": {
                "failureThreshold": 5,
                "httpGet": {
                  "path": "/healthcheck/kubedns",
                  "port": 10054,
                  "scheme": "HTTP"
                },
                "initialDelaySeconds": 60,
                "periodSeconds": 10,
                "successThreshold": 1,
                "timeoutSeconds": 5
              },
              "name": "kubedns",
              "ports": [
                {
                  "containerPort": 10053,
                  "name": "dns-local",
                  "protocol": "UDP"
                },
                {
                  "containerPort": 10053,
                  "name": "dns-tcp-local",
                  "protocol": "TCP"
                },
                {
                  "containerPort": 10055,
                  "name": "metrics",
                  "protocol": "TCP"
                }
              ],
              "readinessProbe": {
                "failureThreshold": 3,
                "httpGet": {
                  "path": "/readiness",
                  "port": 8081,
                  "scheme": "HTTP"
                },
                "initialDelaySeconds": 3,
                "periodSeconds": 10,
                "successThreshold": 1,
                "timeoutSeconds": 5
              },
              "resources": {
                "limits": {
                  "memory": "210Mi"
                },
                "requests": {
                  "cpu": "100m",
                  "memory": "70Mi"
                }
              },
              "securityContext": {
                "allowPrivilegeEscalation": false,
                "readOnlyRootFilesystem": true,
                "runAsGroup": 1001,
                "runAsUser": 1001
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File",
              "volumeMounts": [
                {
                  "mountPath": "/kube-dns-config",
                  "name": "kube-dns-config"
                }
              ]
            },
            {
              "args": [
                "-v=2",
                "-logtostderr",
                "-configDir=/etc/k8s/dns/dnsmasq-nanny",
                "-restartDnsmasq=true",
                "--",
                "-k",
                "--cache-size=1000",
                "--no-negcache",
                "--dns-forward-max=1500",
                "--log-facility=-",
                "--server=/cluster.local/127.0.0.1#10053",
                "--server=/in-addr.arpa/127.0.0.1#10053",
                "--server=/ip6.arpa/127.0.0.1#10053",
                "--max-ttl=30",
                "--max-cache-ttl=30"
              ],
              "image": "gke.gcr.io/k8s-dns-dnsmasq-nanny:1.22.12-gke.0",
              "imagePullPolicy": "IfNotPresent",
              "livenessProbe": {
                "failureThreshold": 5,
                "httpGet": {
                  "path": "/healthcheck/dnsmasq",
                  "port": 10054,
                  "scheme": "HTTP"
                },
                "initialDelaySeconds": 60,
                "periodSeconds": 10,
                "successThreshold": 1,
                "timeoutSeconds": 5
              },
              "name": "dnsmasq",
              "ports": [
                {
                  "containerPort": 53,
                  "name": "dns",
                  "protocol": "UDP"
                },
                {
                  "containerPort": 53,
                  "name": "dns-tcp",
                  "protocol": "TCP"
                }
              ],
              "resources": {
                "requests": {
                  "cpu": "150m",
                  "memory": "20Mi"
                }
              },
              "securityContext": {
                "capabilities": {
                  "add": [
                    "NET_BIND_SERVICE",
                    "SETGID"
                  ],
                  "drop": [
                    "all"
                  ]
                }
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File",
              "volumeMounts": [
                {
                  "mountPath": "/etc/k8s/dns/dnsmasq-nanny",
                  "name": "kube-dns-config"
                }
              ]
            },
            {
              "args": [
                "--v=2",
                "--logtostderr",
                "--probe=kubedns,127.0.0.1:10053,kubernetes.default.svc.cluster.local,5,SRV",
                "--probe=dnsmasq,127.0.0.1:53,kubernetes.default.svc.cluster.local,5,SRV"
              ],
              "image": "gke.gcr.io/k8s-dns-sidecar:1.22.12-gke.0",
              "imagePullPolicy": "IfNotPresent",
              "livenessProbe": {
                "failureThreshold": 5,
                "httpGet": {
                  "path": "/metrics",
                  "port": 10054,
                  "scheme": "HTTP"
                },
                "initialDelaySeconds": 60,
                "periodSeconds": 10,
                "successThreshold": 1,
                "timeoutSeconds": 5
              },
              "name": "sidecar",
              "ports": [
                {
                  "containerPort": 10054,
                  "name": "metrics",
                  "protocol": "TCP"
                }
              ],
              "resources": {
                "requests": {
                  "cpu": "10m",
                  "memory": "20Mi"
                }
              },
              "securityContext": {
                "allowPrivilegeEscalation": false,
                "readOnlyRootFilesystem": true,
                "runAsGroup": 1001,
                "runAsUser": 1001
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File"
            },
            {
              "command": [
                "/monitor",
                "--source=kubedns:http://localhost:10054?whitelisted=probe_kubedns_latency_ms,probe_kubedns_errors,probe_dnsmasq_latency_ms,probe_dnsmasq_errors,dnsmasq_misses,dnsmasq_hits",
                "--stackdriver-prefix=container.googleapis.com/internal/addons",
                "--api-override=https://monitoring.googleapis.com/",
                "--pod-id=$(POD_NAME)",
                "--namespace-id=$(POD_NAMESPACE)",
                "--v=2"
              ],
              "env": [
                {
                  "name": "POD_NAME",
                  "valueFrom": {
                    "fieldRef": {
                      "apiVersion": "v1",
                      "fieldPath": "metadata.name"
                    }
                  }
                },
                {
                  "name": "POD_NAMESPACE",
                  "valueFrom": {
                    "fieldRef": {
                      "apiVersion": "v1",
                      "fieldPath": "metadata.namespace"
                    }
                  }
                }
              ],
              "image": "gke.gcr.io/prometheus-to-sd:v0.11.3-gke.0",
              "imagePullPolicy": "IfNotPresent",
              "name": "prometheus-to-sd",
              "resources": {},
              "securityContext": {
                "allowPrivilegeEscalation": false,
                "readOnlyRootFilesystem": true,
                "runAsGroup": 1001,
                "runAsUser": 1001
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File"
            }
          ],
          "dnsPolicy": "Default",
          "nodeSelector": {
            "kubernetes.io/os": "linux"
          },
          "priorityClassName": "system-cluster-critical",
          "restartPolicy": "Always",
          "schedulerName": "default-scheduler",
          "securityContext": {
            "fsGroup": 65534,
            "supplementalGroups": [
              65534
            ]
          },
          "serviceAccount": "kube-dns",
          "serviceAccountName": "kube-dns",
          "terminationGracePeriodSeconds": 30,
          "tolerations": [
            {
              "key": "CriticalAddonsOnly",
              "operator": "Exists"
            },
            {
              "key": "components.gke.io/gke-managed-components",
              "operator": "Exists"
            },
            {
              "effect": "NoSchedule",
              "key": "kubernetes.io/arch",
              "operator": "Equal",
              "value": "arm64"
            }
          ],
          "volumes": [
            {
              "configMap": {
                "defaultMode": 420,
                "name": "kube-dns",
                "optional": true
              },
              "name": "kube-dns-config"
            }
          ]
        }
      }
    },
    "status": {
      "availableReplicas": 2,
      "conditions": [
        {
          "lastTransitionTime": "2022-02-15T05:46:18Z",
          "lastUpdateTime": "2023-02-10T03:34:47Z",
          "message": "ReplicaSet \"kube-dns-698cf6b7dc\" has successfully progressed.",
          "reason": "NewReplicaSetAvailable",
          "status": "True",
          "type": "Progressing"
        },
        {
          "lastTransitionTime": "2023-02-26T05:13:07Z",
          "lastUpdateTime": "2023-02-26T05:13:07Z",
          "message": "Deployment has minimum availability.",
          "reason": "MinimumReplicasAvailable",
          "status": "True",
          "type": "Available"
        }
      ],
      "observedGeneration": 9,
      "readyReplicas": 2,
      "replicas": 2,
      "updatedReplicas": 2
    }
  },
  {
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
      "annotations": {
        "deployment.kubernetes.io/revision": "2"
      },
      "creationTimestamp": "2022-02-15T05:46:17Z",
      "generation": 3,
      "labels": {
        "addonmanager.kubernetes.io/mode": "Reconcile",
        "k8s-app": "kube-dns-autoscaler",
        "kubernetes.io/cluster-service": "true"
      },
      "name": "kube-dns-autoscaler",
      "namespace": "kube-system",
      "resourceVersion": "284150211",
      "uid": "4fea76ed-4820-4cff-8e3b-658ff52d97cd"
    },
    "spec": {
      "progressDeadlineSeconds": 600,
      "replicas": 1,
      "revisionHistoryLimit": 10,
      "selector": {
        "matchLabels": {
          "k8s-app": "kube-dns-autoscaler"
        }
      },
      "strategy": {
        "rollingUpdate": {
          "maxSurge": "25%",
          "maxUnavailable": "25%"
        },
        "type": "RollingUpdate"
      },
      "template": {
        "metadata": {
          "creationTimestamp": null,
          "labels": {
            "k8s-app": "kube-dns-autoscaler"
          }
        },
        "spec": {
          "containers": [
            {
              "command": [
                "/cluster-proportional-autoscaler",
                "--namespace=kube-system",
                "--configmap=kube-dns-autoscaler",
                "--target=Deployment/kube-dns",
                "--default-params={\"linear\":{\"coresPerReplica\":256,\"nodesPerReplica\":16,\"preventSinglePointFailure\":true,\"includeUnschedulableNodes\":true}}",
                "--logtostderr=true",
                "--v=2"
              ],
              "image": "gke.gcr.io/cluster-proportional-autoscaler:1.8.4-gke.1",
              "imagePullPolicy": "IfNotPresent",
              "name": "autoscaler",
              "resources": {
                "requests": {
                  "cpu": "20m",
                  "memory": "10Mi"
                }
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File"
            }
          ],
          "dnsPolicy": "ClusterFirst",
          "nodeSelector": {
            "kubernetes.io/os": "linux"
          },
          "priorityClassName": "system-cluster-critical",
          "restartPolicy": "Always",
          "schedulerName": "default-scheduler",
          "securityContext": {
            "fsGroup": 65534,
            "seccompProfile": {
              "type": "RuntimeDefault"
            },
            "supplementalGroups": [
              65534
            ]
          },
          "serviceAccount": "kube-dns-autoscaler",
          "serviceAccountName": "kube-dns-autoscaler",
          "terminationGracePeriodSeconds": 30,
          "tolerations": [
            {
              "key": "CriticalAddonsOnly",
              "operator": "Exists"
            },
            {
              "key": "components.gke.io/gke-managed-components",
              "operator": "Exists"
            }
          ]
        }
      }
    },
    "status": {
      "availableReplicas": 1,
      "conditions": [
        {
          "lastTransitionTime": "2022-02-15T05:46:18Z",
          "lastUpdateTime": "2022-07-14T16:27:53Z",
          "message": "ReplicaSet \"kube-dns-autoscaler-f4d55555\" has successfully progressed.",
          "reason": "NewReplicaSetAvailable",
          "status": "True",
          "type": "Progressing"
        },
        {
          "lastTransitionTime": "2023-02-25T05:13:04Z",
          "lastUpdateTime": "2023-02-25T05:13:04Z",
          "message": "Deployment has minimum availability.",
          "reason": "MinimumReplicasAvailable",
          "status": "True",
          "type": "Available"
        }
      ],
      "observedGeneration": 3,
      "readyReplicas": 1,
      "replicas": 1,
      "updatedReplicas": 1
    }
  },
  {
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
      "annotations": {
        "components.gke.io/layer": "addon",
        "deployment.kubernetes.io/revision": "4",
        "seccomp.security.alpha.kubernetes.io/pod": "runtime/default"
      },
      "creationTimestamp": "2022-02-15T05:46:19Z",
      "generation": 5,
      "labels": {
        "addonmanager.kubernetes.io/mode": "Reconcile",
        "k8s-app": "glbc",
        "kubernetes.io/cluster-service": "true",
        "kubernetes.io/name": "GLBC"
      },
      "name": "l7-default-backend",
      "namespace": "kube-system",
      "resourceVersion": "284150849",
      "uid": "6cf80b6a-64f3-4228-bb77-20888278b81e"
    },
    "spec": {
      "progressDeadlineSeconds": 600,
      "replicas": 1,
      "revisionHistoryLimit": 10,
      "selector": {
        "matchLabels": {
          "k8s-app": "glbc"
        }
      },
      "strategy": {
        "rollingUpdate": {
          "maxSurge": "25%",
          "maxUnavailable": "25%"
        },
        "type": "RollingUpdate"
      },
      "template": {
        "metadata": {
          "annotations": {
            "components.gke.io/component-name": "l7-lb-controller-combined",
            "components.gke.io/component-version": "1.14.8-gke.0",
            "seccomp.security.alpha.kubernetes.io/pod": "runtime/default"
          },
          "creationTimestamp": null,
          "labels": {
            "k8s-app": "glbc",
            "name": "glbc"
          }
        },
        "spec": {
          "containers": [
            {
              "image": "gke.gcr.io/ingress-gce-404-server-with-metrics:v1.13.4",
              "imagePullPolicy": "IfNotPresent",
              "livenessProbe": {
                "failureThreshold": 3,
                "httpGet": {
                  "path": "/healthz",
                  "port": 8080,
                  "scheme": "HTTP"
                },
                "initialDelaySeconds": 30,
                "periodSeconds": 10,
                "successThreshold": 1,
                "timeoutSeconds": 5
              },
              "name": "default-http-backend",
              "ports": [
                {
                  "containerPort": 8080,
                  "protocol": "TCP"
                }
              ],
              "resources": {
                "requests": {
                  "cpu": "10m",
                  "memory": "20Mi"
                }
              },
              "securityContext": {
                "allowPrivilegeEscalation": false,
                "capabilities": {
                  "drop": [
                    "all"
                  ]
                },
                "runAsGroup": 1000,
                "runAsUser": 1000
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File"
            }
          ],
          "dnsPolicy": "ClusterFirst",
          "restartPolicy": "Always",
          "schedulerName": "default-scheduler",
          "securityContext": {},
          "terminationGracePeriodSeconds": 30,
          "tolerations": [
            {
              "key": "components.gke.io/gke-managed-components",
              "operator": "Exists"
            }
          ]
        }
      }
    },
    "status": {
      "availableReplicas": 1,
      "conditions": [
        {
          "lastTransitionTime": "2022-02-15T05:46:19Z",
          "lastUpdateTime": "2023-01-14T05:21:32Z",
          "message": "ReplicaSet \"l7-default-backend-7dc577646d\" has successfully progressed.",
          "reason": "NewReplicaSetAvailable",
          "status": "True",
          "type": "Progressing"
        },
        {
          "lastTransitionTime": "2023-02-25T05:14:09Z",
          "lastUpdateTime": "2023-02-25T05:14:09Z",
          "message": "Deployment has minimum availability.",
          "reason": "MinimumReplicasAvailable",
          "status": "True",
          "type": "Available"
        }
      ],
      "observedGeneration": 5,
      "readyReplicas": 1,
      "replicas": 1,
      "updatedReplicas": 1
    }
  },
  {
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
      "annotations": {
        "deployment.kubernetes.io/revision": "9"
      },
      "creationTimestamp": "2022-03-15T15:56:43Z",
      "generation": 10,
      "labels": {
        "addonmanager.kubernetes.io/mode": "Reconcile",
        "k8s-app": "metrics-server",
        "version": "v0.4.5"
      },
      "name": "metrics-server-v0.4.5",
      "namespace": "kube-system",
      "resourceVersion": "305781911",
      "uid": "ff19f9a9-df5b-4ed9-9d92-674b2a598c87"
    },
    "spec": {
      "progressDeadlineSeconds": 600,
      "replicas": 1,
      "revisionHistoryLimit": 10,
      "selector": {
        "matchLabels": {
          "k8s-app": "metrics-server",
          "version": "v0.4.5"
        }
      },
      "strategy": {
        "rollingUpdate": {
          "maxSurge": "25%",
          "maxUnavailable": "25%"
        },
        "type": "RollingUpdate"
      },
      "template": {
        "metadata": {
          "creationTimestamp": null,
          "labels": {
            "k8s-app": "metrics-server",
            "version": "v0.4.5"
          },
          "name": "metrics-server"
        },
        "spec": {
          "containers": [
            {
              "command": [
                "/metrics-server",
                "--metric-resolution=30s",
                "--kubelet-port=10255",
                "--deprecated-kubelet-completely-insecure=true",
                "--kubelet-preferred-address-types=InternalIP,Hostname,InternalDNS,ExternalDNS,ExternalIP",
                "--cert-dir=/tmp",
                "--secure-port=10250"
              ],
              "image": "gke.gcr.io/metrics-server-amd64:v0.4.5-gke.0",
              "imagePullPolicy": "IfNotPresent",
              "name": "metrics-server",
              "ports": [
                {
                  "containerPort": 10250,
                  "name": "https",
                  "protocol": "TCP"
                }
              ],
              "resources": {
                "limits": {
                  "cpu": "43m",
                  "memory": "55Mi"
                },
                "requests": {
                  "cpu": "43m",
                  "memory": "55Mi"
                }
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File",
              "volumeMounts": [
                {
                  "mountPath": "/tmp",
                  "name": "tmp-dir"
                }
              ]
            },
            {
              "command": [
                "/pod_nanny",
                "--config-dir=/etc/config",
                "--cpu=40m",
                "--extra-cpu=0.5m",
                "--memory=35Mi",
                "--extra-memory=4Mi",
                "--threshold=5",
                "--deployment=metrics-server-v0.4.5",
                "--container=metrics-server",
                "--poll-period=30000",
                "--estimator=exponential",
                "--scale-down-delay=24h",
                "--minClusterSize=5",
                "--use-metrics=true"
              ],
              "env": [
                {
                  "name": "MY_POD_NAME",
                  "valueFrom": {
                    "fieldRef": {
                      "apiVersion": "v1",
                      "fieldPath": "metadata.name"
                    }
                  }
                },
                {
                  "name": "MY_POD_NAMESPACE",
                  "valueFrom": {
                    "fieldRef": {
                      "apiVersion": "v1",
                      "fieldPath": "metadata.namespace"
                    }
                  }
                }
              ],
              "image": "gke.gcr.io/addon-resizer:1.8.13-gke.0",
              "imagePullPolicy": "IfNotPresent",
              "name": "metrics-server-nanny",
              "resources": {
                "limits": {
                  "cpu": "100m",
                  "memory": "300Mi"
                },
                "requests": {
                  "cpu": "5m",
                  "memory": "50Mi"
                }
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File",
              "volumeMounts": [
                {
                  "mountPath": "/etc/config",
                  "name": "metrics-server-config-volume"
                }
              ]
            }
          ],
          "dnsPolicy": "ClusterFirst",
          "nodeSelector": {
            "kubernetes.io/os": "linux"
          },
          "priorityClassName": "system-cluster-critical",
          "restartPolicy": "Always",
          "schedulerName": "default-scheduler",
          "securityContext": {
            "seccompProfile": {
              "type": "RuntimeDefault"
            }
          },
          "serviceAccount": "metrics-server",
          "serviceAccountName": "metrics-server",
          "terminationGracePeriodSeconds": 30,
          "tolerations": [
            {
              "key": "CriticalAddonsOnly",
              "operator": "Exists"
            },
            {
              "key": "components.gke.io/gke-managed-components",
              "operator": "Exists"
            }
          ],
          "volumes": [
            {
              "configMap": {
                "defaultMode": 420,
                "name": "metrics-server-config"
              },
              "name": "metrics-server-config-volume"
            },
            {
              "emptyDir": {},
              "name": "tmp-dir"
            }
          ]
        }
      }
    },
    "status": {
      "availableReplicas": 1,
      "conditions": [
        {
          "lastTransitionTime": "2022-03-15T15:56:43Z",
          "lastUpdateTime": "2022-08-06T11:19:36Z",
          "message": "ReplicaSet \"metrics-server-v0.4.5-fb4c49dd6\" has successfully progressed.",
          "reason": "NewReplicaSetAvailable",
          "status": "True",
          "type": "Progressing"
        },
        {
          "lastTransitionTime": "2023-02-25T05:13:16Z",
          "lastUpdateTime": "2023-02-25T05:13:16Z",
          "message": "Deployment has minimum availability.",
          "reason": "MinimumReplicasAvailable",
          "status": "True",
          "type": "Available"
        }
      ],
      "observedGeneration": 10,
      "readyReplicas": 1,
      "replicas": 1,
      "updatedReplicas": 1
    }
  }
]`
