package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println(HelloWorld())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", indexHandler)
	log.Printf("listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// HelloWorld is a function that returns a string containing "hello world".
func HelloWorld() string {
	return "hello world"
}

// indexHandler serves the blue fluid UI page.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, indexHTML)
}

const indexHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>Hello World</title>
<style>
  html, body { margin: 0; height: 100%; overflow: hidden; }
  body {
    display: flex; align-items: center; justify-content: center;
    background: linear-gradient(160deg, #021a3a 0%, #04336e 50%, #0a5fb4 100%);
    font-family: system-ui, -apple-system, "Segoe UI", Roboto, sans-serif;
    color: #e6f2ff;
  }
  .blob {
    position: absolute; width: 60vmax; height: 60vmax; border-radius: 50%;
    filter: blur(60px); opacity: 0.55; mix-blend-mode: screen;
    animation: drift 18s ease-in-out infinite alternate;
  }
  .b1 { background: #1e90ff; top: -20%; left: -15%; }
  .b2 { background: #00c6ff; bottom: -25%; right: -10%; animation-duration: 22s; }
  .b3 { background: #3a6bff; top: 20%; right: 10%; width: 45vmax; height: 45vmax; animation-duration: 26s; }
  @keyframes drift {
    0%   { transform: translate(0, 0) scale(1); }
    50%  { transform: translate(8vw, 6vh) scale(1.15); }
    100% { transform: translate(-6vw, 10vh) scale(0.9); }
  }
  .waves { position: absolute; left: 0; bottom: 0; width: 200%; height: 35vh; }
  .waves use { animation: wave 9s linear infinite; }
  .waves use:nth-child(2) { animation-duration: 13s; animation-direction: reverse; }
  .waves use:nth-child(3) { animation-duration: 17s; }
  @keyframes wave {
    from { transform: translateX(0); }
    to   { transform: translateX(-50%); }
  }
  h1 {
    position: relative; margin: 0; font-size: clamp(2.5rem, 8vw, 6rem);
    font-weight: 700; letter-spacing: 0.04em;
    text-shadow: 0 0 25px rgba(0, 170, 255, 0.8);
  }
</style>
</head>
<body>
  <div class="blob b1"></div>
  <div class="blob b2"></div>
  <div class="blob b3"></div>
  <svg class="waves" viewBox="0 0 2880 320" preserveAspectRatio="none" aria-hidden="true">
    <defs>
      <path id="w" d="M0,160 C320,260 640,60 960,160 C1280,260 1600,60 1920,160 C2240,260 2560,60 2880,160 L2880,320 L0,320 Z"/>
    </defs>
    <use href="#w" fill="rgba(30,144,255,0.35)"/>
    <use href="#w" fill="rgba(0,198,255,0.30)" y="20"/>
    <use href="#w" fill="rgba(58,107,255,0.45)" y="40"/>
  </svg>
  <h1>hello world</h1>
</body>
</html>
`
