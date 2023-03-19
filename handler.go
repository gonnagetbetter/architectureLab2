package lab2

import "io"

type ComputeHandler struct {
  Input  io.Reader
  Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
  s := make([]byte, 0)
  p := make([]byte, 8)

  for {
    n, err := ch.Input.Read(p)
    s = append(s, p[:n]...)
    if err == io.EOF {
      break
    }
  }

  result, err := PostfixToPrefix(string(s))

  if err != nil {
    return err
  }
  
  ch.Output.Write([]byte(result))
  return nil
}
