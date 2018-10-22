# Workers serve on heavy tasks
## Situation
- high traffic requests
- heavy work for each request
## Solution
- separate heavy works with response (assumsion: able to separate heavy tasks to have quick response)
- tasks queue and worker pool are created to handle number of go-routines generated 
![Worker serves heavy task](https://preview.ibb.co/emChvf/Workers-Serve-Heavy-Task.jpg "Workers serve heavy tasks")
