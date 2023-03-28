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

var PARSED_JSON = `[{"apiVersion":"apps/v1","kind":"Deployment","metadata":{"annotations":{"argocd.argoproj.io/sync-wave":"3","cnrm.cloud.google.com/project-id":"test-prod-customers","deployment.kubernetes.io/revision":"129","kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"apps/v1\",\"kind\":\"Deployment\",\"metadata\":{\"annotations\":{\"argocd.argoproj.io/sync-wave\":\"3\",\"cnrm.cloud.google.com/project-id\":\"test-prod-customers\"},\"labels\":{\"app.kubernetes.io/component\":\"test-product-deployment\",\"app.kubernetes.io/instance\":\"prod-customer-data\",\"app.kubernetes.io/name\":\"customer-data\",\"app.kubernetes.io/part-of\":\"events\",\"k8s.test.com/team\":\"data-platform\",\"tags.datadoghq.com/env\":\"prod\",\"tags.datadoghq.com/service\":\"customer-data\",\"tags.datadoghq.com/version\":\"56869a5\"},\"name\":\"customer-data\",\"namespace\":\"customer-data\"},\"spec\":{\"selector\":{\"matchLabels\":{\"app\":\"customer-data\",\"app.kubernetes.io/part-of\":\"events\",\"k8s.test.com/team\":\"data-platform\",\"tags.datadoghq.com/env\":\"prod\"}},\"strategy\":{\"rollingUpdate\":{\"maxSurge\":\"10%\",\"maxUnavailable\":\"0%\"},\"type\":\"RollingUpdate\"},\"template\":{\"metadata\":{\"annotations\":{\"cnrm.cloud.google.com/project-id\":\"test-prod-customers\",\"prometheus.io/port\":\"9102\",\"prometheus.io/scrape\":\"true\",\"sidecar.istio.io/capNetBindService\":\"true\",\"sidecar.istio.io/userVolume\":\"[{\\\"name\\\":\\\"descriptor\\\", \\\"secret\\\":{\\\"secretName\\\":\\\"descriptors-pb\\\"}}]\",\"sidecar.istio.io/userVolumeMount\":\"[{\\\"name\\\":\\\"descriptor\\\", \\\"mountPath\\\":\\\"/etc/envoy\\\", \\\"readonly\\\":true}]\"},\"labels\":{\"app\":\"customer-data\",\"app.kubernetes.io/component\":\"test-product-deployment\",\"app.kubernetes.io/name\":\"customer-data\",\"app.kubernetes.io/part-of\":\"events\",\"k8s.test.com/team\":\"data-platform\",\"tags.datadoghq.com/env\":\"prod\",\"tags.datadoghq.com/service\":\"customer-data\",\"tags.datadoghq.com/version\":\"56869a5\"}},\"spec\":{\"containers\":[{\"args\":[\"server\"],\"env\":[{\"name\":\"SERVICE_NAME\",\"value\":\"customer-data\"},{\"name\":\"HOST_IP\",\"valueFrom\":{\"fieldRef\":{\"fieldPath\":\"status.hostIP\"}}},{\"name\":\"TRACE_ADDRESS\",\"value\":\"$(HOST_IP):4317\"},{\"name\":\"BLACKLIST_FILE\",\"value\":\"/app/blacklist.yaml\"},{\"name\":\"DD_ENV\",\"valueFrom\":{\"fieldRef\":{\"fieldPath\":\"metadata.labels['tags.datadoghq.com/env']\"}}},{\"name\":\"DD_SERVICE\",\"valueFrom\":{\"fieldRef\":{\"fieldPath\":\"metadata.labels['tags.datadoghq.com/service']\"}}},{\"name\":\"DD_VERSION\",\"valueFrom\":{\"fieldRef\":{\"fieldPath\":\"metadata.labels['tags.datadoghq.com/version']\"}}},{\"name\":\"GARBAGE\",\"value\":\"delete-me-later\"}],\"envFrom\":[{\"configMapRef\":{\"name\":\"config-86t98h48c7\"}},{\"configMapRef\":{\"name\":\"rpc-config-ht9m76d66k\"}},{\"secretRef\":{\"name\":\"customer-data-secrets\"}}],\"image\":\"us.gcr.io/test-ops/apps/customer-data:56869a5\",\"imagePullPolicy\":\"Always\",\"livenessProbe\":{\"exec\":{\"command\":[\"/bin/grpc_health_probe\",\"-addr=:50051\"]},\"initialDelaySeconds\":10,\"periodSeconds\":10},\"name\":\"customer-data\",\"ports\":[{\"containerPort\":50051,\"name\":\"grpc\"},{\"containerPort\":8080,\"name\":\"http\"},{\"containerPort\":9102,\"name\":\"http-metrics\"}],\"readinessProbe\":{\"exec\":{\"command\":[\"/bin/grpc_health_probe\",\"-addr=:50051\"]},\"initialDelaySeconds\":5},\"resources\":{\"limits\":{\"memory\":\"2Gi\"},\"requests\":{\"cpu\":\"2\",\"memory\":\"1Gi\"}},\"volumeMounts\":[{\"mountPath\":\"/app\",\"name\":\"blacklist\"}]}],\"dnsConfig\":{\"options\":[{\"name\":\"ndots\",\"value\":\"3\"}]},\"serviceAccountName\":\"argo-sa-customer-data\",\"volumes\":[{\"configMap\":{\"name\":\"blacklist\"},\"name\":\"blacklist\"}]}}}}\n"},"creationTimestamp":"2021-08-24T19:06:44Z","generation":7204,"labels":{"app.kubernetes.io/component":"test-product-deployment","app.kubernetes.io/instance":"prod-customer-data","app.kubernetes.io/name":"customer-data","app.kubernetes.io/part-of":"events","k8s.test.com/team":"data-platform","tags.datadoghq.com/env":"prod","tags.datadoghq.com/service":"customer-data","tags.datadoghq.com/version":"56869a5"},"name":"customer-data","namespace":"customer-data","resourceVersion":"4020930045","uid":"09026a6a-3310-4623-8ce2-f1dd310c9391"},"spec":{"progressDeadlineSeconds":600,"replicas":10,"revisionHistoryLimit":10,"selector":{"matchLabels":{"app":"customer-data","app.kubernetes.io/part-of":"events","k8s.test.com/team":"data-platform","tags.datadoghq.com/env":"prod"}},"strategy":{"rollingUpdate":{"maxSurge":"10%","maxUnavailable":"0%"},"type":"RollingUpdate"},"template":{"metadata":{"annotations":{"cnrm.cloud.google.com/project-id":"test-prod-customers","kubectl.kubernetes.io/restartedAt":"2022-09-06T20:46:32-04:00","prometheus.io/port":"9102","prometheus.io/scrape":"true","sidecar.istio.io/capNetBindService":"true","sidecar.istio.io/userVolume":"[{\"name\":\"descriptor\", \"secret\":{\"secretName\":\"descriptors-pb\"}}]","sidecar.istio.io/userVolumeMount":"[{\"name\":\"descriptor\", \"mountPath\":\"/etc/envoy\", \"readonly\":true}]"},"creationTimestamp":null,"labels":{"app":"customer-data","app.kubernetes.io/component":"test-product-deployment","app.kubernetes.io/name":"customer-data","app.kubernetes.io/part-of":"events","k8s.test.com/team":"data-platform","tags.datadoghq.com/env":"prod","tags.datadoghq.com/service":"customer-data","tags.datadoghq.com/version":"56869a5"}},"spec":{"containers":[{"args":["server"],"env":[{"name":"SERVICE_NAME","value":"customer-data"},{"name":"HOST_IP","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"status.hostIP"}}},{"name":"TRACE_ADDRESS","value":"$(HOST_IP):4317"},{"name":"BLACKLIST_FILE","value":"/app/blacklist.yaml"},{"name":"DD_ENV","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"metadata.labels['tags.datadoghq.com/env']"}}},{"name":"DD_SERVICE","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"metadata.labels['tags.datadoghq.com/service']"}}},{"name":"DD_VERSION","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"metadata.labels['tags.datadoghq.com/version']"}}},{"name":"GARBAGE","value":"delete-me-later"}],"envFrom":[{"configMapRef":{"name":"config-86t98h48c7"}},{"configMapRef":{"name":"rpc-config-ht9m76d66k"}},{"secretRef":{"name":"customer-data-secrets"}}],"image":"us.gcr.io/test-ops/apps/customer-data:56869a5","imagePullPolicy":"Always","livenessProbe":{"exec":{"command":["/bin/grpc_health_probe","-addr=:50051"]},"failureThreshold":3,"initialDelaySeconds":10,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":1},"name":"customer-data","ports":[{"containerPort":50051,"name":"grpc","protocol":"TCP"},{"containerPort":8080,"name":"http","protocol":"TCP"},{"containerPort":9102,"name":"http-metrics","protocol":"TCP"}],"readinessProbe":{"exec":{"command":["/bin/grpc_health_probe","-addr=:50051"]},"failureThreshold":3,"initialDelaySeconds":5,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":1},"resources":{"limits":{"memory":"2Gi"},"requests":{"cpu":"2","memory":"1Gi"}},"terminationMessagePath":"/dev/termination-log","terminationMessagePolicy":"File","volumeMounts":[{"mountPath":"/app","name":"blacklist"}]}],"dnsConfig":{"options":[{"name":"ndots","value":"3"}]},"dnsPolicy":"ClusterFirst","restartPolicy":"Always","schedulerName":"default-scheduler","securityContext":{},"serviceAccount":"argo-sa-customer-data","serviceAccountName":"argo-sa-customer-data","terminationGracePeriodSeconds":30,"volumes":[{"configMap":{"defaultMode":420,"name":"blacklist"},"name":"blacklist"}]}}},"status":{"availableReplicas":10,"conditions":[{"lastTransitionTime":"2021-08-24T19:40:12Z","lastUpdateTime":"2022-11-11T21:40:09Z","message":"ReplicaSet \"customer-data-544fb5d4c6\" has successfully progressed.","reason":"NewReplicaSetAvailable","status":"True","type":"Progressing"},{"lastTransitionTime":"2022-11-15T15:09:29Z","lastUpdateTime":"2022-11-15T15:09:29Z","message":"Deployment has minimum availability.","reason":"MinimumReplicasAvailable","status":"True","type":"Available"}],"observedGeneration":7204,"readyReplicas":10,"replicas":10,"updatedReplicas":10}}]`
