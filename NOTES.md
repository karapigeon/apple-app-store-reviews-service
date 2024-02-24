# Notes

* My current machine is very lacking so using a GitHub Codespace to work on this.
* I was originally going to write the back-end in Rust but then going through the document, it mentioned primarily only using standard library APIs and the Rust standard library is very low level which would make a lot of "basic" functionality more involved.
* Given this, I'm going to write the back-end in Go because the Go standard library is very good and it would be the same amount of _floundering_ than if I were to use a different language. I've worked with TypeScript, JavaScript, Go, and Rust in my last couple roles so might as well use the stack preferred by the exercise.
* I initially worked out of a single file and then I refactored it out into multiple after initial implementation functioned.
* Back-end is functioning and saves the dataset to a `595068606.json` file on disk.
* I originally used [`npx create-next-app@latest`](https://react.dev/learn/start-a-new-react-project#nextjs-pages-router) to create the React app however then I ran into issues with Next being able to access the local file due to server component limitations and as mentioned in our call, the specifics of front-end frameworks is not my strongest area so through troubleshooting this, I found that Remix was a better framework/DX for this type of app and used [`npx create-remix`](https://react.dev/learn/start-a-new-react-project#remix) to get started again.
* Next up is the front-end scaffold to access the data. 