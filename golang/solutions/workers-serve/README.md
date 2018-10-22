# Worker serve on heavy task
## Situation
- high traffic requests
- heavy work for each request
## Solution
- separate heavy work with response (assumsion: able to separate heavy task to have quick response)
- tasks queue and worker pool are created to handle number of go-routines generated 
![Worker serves heavy task](https://preview.ibb.co/emChvf/Workers-Serve-Heavy-Task.jpg "Worker serves heavy task")
