package main

import (
	"bytes"
	"errors"
	"github.com/panjf2000/ants"
	"io"
	"log"
	"sync"
	"time"
	"os"
	"math/rand"
	"fmt"
)

type Runner struct {
	tasks []func(int)
	complete chan error
	timeout<-chan time.Time
	interrput chan os.Signal
}

func New(tm time.Duration) *Runner {
     return &Runner{
     	complete:make(chan error),
     	timeout:time.After(tm),
     	interrput:make(chan os.Signal,1),


	 }
}

func (r *Runner)Add(task ...func(int))  {
	r.tasks = append(r.tasks,task...)
}
func generate_Random() int{
	rand.Seed(time.Now().Unix())

	rnd := rand.Intn(100)

	fmt.Printf("rand is %v\n",rnd)

	return rnd
}

type pool struct {
	m sync.Mutex
	res chan io.Closer
	factory func()(io.Closer,error)
	closed bool
}

func New2(fn func()(io.Closer,error),size int)(*pool,error)  {
	if size <= 0{
		return nil,errors.New("size的值太小了。")
	}
	return &pool{
		factory:fn,
		res: make(chan io.Closer,size),

	},nil
}

func (p *pool) Acquire()(io.Closer,error)  {
	select {
	case r,ok := <-p.res:
		log.Println("Acquire: 共享资源")
	if !ok{
		return nil,ants.ErrPoolClosed
	}
	return r,nil
	default:
		log.Println("Acquire:新生成资源")
	return p.factory()
	}
}

func(p *pool) Close()  {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed{
		return
	}
	p.closed = true

	close(p.res)

	for r:=range p.res{
		r.Close()
	}

}
func main()  {

//	generate_Random()

//	fmt.Println("Hello Go!")


	var b bytes.Buffer

	b.Write([]byte("你好"))

	fmt.Fprint(&b,",","http://www.flysnow.org")

	b.WriteTo(os.Stdout)























}
