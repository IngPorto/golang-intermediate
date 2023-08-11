package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Job representa una tarea con su nombre, retraso y número asociado.
type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

// Worker representa un trabajador que puede ejecutar tareas.
type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

// Dispatcher coordina la asignación de tareas a los trabajadores.
type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

// NewWorker crea y devuelve un nuevo trabajador con el id proporcionado y el pool de trabajadores dado.
func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool),
	}
}

// Start inicia la rutina de trabajo del trabajador.
func (w Worker) Start() {
	go func() {
		for {
			// Registra el canal de la cola de trabajos del trabajador en el pool de trabajadores.
			w.WorkerPool <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				// Ejecuta la tarea y registra su progreso.
				fmt.Printf("Worker %d started job %s\n", w.Id, job.Name)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker %d finished job %s and fib is %d\n", w.Id, job.Name, fib)
			case <-w.QuitChan:
				// Detiene la rutina del trabajador y muestra un mensaje de parada.
				fmt.Printf("Worker %d stopped\n", w.Id)
				return
			}
		}
	}()
}

// Stop detiene la rutina del trabajador.
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

// Fibonacci calcula el número Fibonacci recursivamente.
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// NewDispatcher crea y devuelve un nuevo despachador con la cola de trabajos y el número máximo de trabajadores.
func NewDispatcher(JobQueue chan Job, MaxWorkers int) *Dispatcher {
	worker := make(chan chan Job, MaxWorkers)
	return &Dispatcher{
		JobQueue:   JobQueue,
		MaxWorkers: MaxWorkers,
		WorkerPool: worker,
	}
}

// Dispatch inicia la asignación de trabajos a los trabajadores.
func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			// Asigna el trabajo a un trabajador disponible del pool.
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}


//--------------------
func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" { // GET, PUT, DELETE
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Invalid Delay" + err.Error(), http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid Value"+ err.Error(), http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Invalid Name", http.StatusBadRequest)
		return
	}

	job := Job{
		Name:   name,
		Delay:  delay,
		Number: value,
	}

	jobQueue <- job // ESTO INICIALIZA TODO EL PROCESO PARA QUE LOS WORKERS SE PONGAN EN ACCION

	w.WriteHeader(http.StatusOK)
}

func main() {
	const (
		maxWorkers = 4
		maxQueueSize = 20
		port = ":8081"
	)

	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)

	dispatcher.Run()

	// http://localhost:8081/fib
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})

	// Ocurre si la función deja de funcionar, podemos ver el error con log.Fatal
	// sino estuviera encerrado en esa función, se bloquería el programa
	log.Fatal(http.ListenAndServe(port, nil))

}