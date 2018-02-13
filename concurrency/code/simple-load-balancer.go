func Run() {
   in, out := make(chan *Work), make(chan *Work)
   for i := 0; i < NumWorkers; i++ {
       go worker(in, out)
   }
   go sendLotsOfWork(in)
   receiveLotsOfResults(out)
}
