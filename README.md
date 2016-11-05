# detectlanguage
A simple client for detectlanguage.com in go
- Detect a sentence
- Detect multiple sentences
- Get account's status
- Get list of languages current supported

## How to use

`import github.com/virusvn/detectlanguage

detectlanguage.Is("en", "your-api-key", "Spratlys and Paracels belong to Vietnam") //true
lang := detectlanguage.Detect("your-api-key", "Trường Sa, Hoàng Sa thuộc Việt Nam") //lang.Data.Detections[0].Language: vi
`

