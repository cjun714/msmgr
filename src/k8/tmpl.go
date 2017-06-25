package k8

// SVCTmpl is used as SVC template
const RCTmpl string = `
kind: ReplicationController
apiVersion: v1
metadata:
  name: DemoAPP
  labels:
    name: DemoAPP
spec:
  replicas: 1
  selector:
    name: DemoAPP
  template:
    metadata:
      labels:
        name: DemoAPP
    spec:
      containers:
      - name: DemoAPP
        image: 20.26.33.121:5000/DemoAPP
        resources:
          limits:
            cpu: 600m
            memory: 600Mi
          requests:
            cpu: 500m
            memory: 500Mi
        env:
        - name: SPRING_CONFIG_URI
          value: 'http://ms:Microservice123!@20.26.33.122:32011'
        - name: CONFIG_ACTIVE
          value: "test"
        command: ["sh","-c","java -server -Xms256m -Xmx256m  -Djava.io.tmpdir=/var/tmp -Duser.timezone=Asia/Shanghai -jar DemoAPP.jar --spring.profiles.active=${CONFIG_ACTIVE} --server.port=8080"]
        ports:
        - containerPort: 8080
`

// RCTmpl is used as SVC template
const SVCTmpl string = `
kind: Service
apiVersion: v1
metadata:
  name: DemoAPP
  labels:
    name: DemoAPP
spec:
  type: NodePort
  ports:
  - nodePort: 32112
    targetPort: 8080
    port: 8080
  selector:
    name: eureka-server
`
