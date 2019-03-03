// Sample language-quickstart uses the Google Cloud Natural API to analyze the
// sentiment of "Hello, world!".
package main

import (
        "context"
        "fmt"
        "log"

        language "cloud.google.com/go/language/apiv1"
        languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func main() {
        ctx := context.Background()

        // Creates a client.
        client, err := language.NewClient(ctx)
        if err != nil {
                log.Fatalf("Failed to create client: %v", err)
        }

        // Sets the text to analyze.
        // text := "Hello world!"
        // text := "最近花粉症で目がしょぼしょぼするのか寝不足で目がしょぼしょぼするのかわからない。"
        text := "NGT山口真帆暴行事件を起こした人とかを見ても、電車に乗ってるうるさい人、酔っ払ってる人、下らない話をしている主婦たちを見ていても、この人達とは隔絶された世界で生きていたいなあと思う。" +
                "一方で、東南アジアのスラム街とか日本のドヤ街が好きで旅行のときに立ち寄ってしまう。\n一見、両者は同質に見える。しかし、前者は群れていて縛られている（縛られるのが好きなのだ）、後者は孤独でかつ自由だという違いがある。" +
                "集団に嫌気が差し、自ら一人を選ぶ人が僕は好きだ。そういう人とは個別に積極的に関わっていきたい。\n集団にいると自由が奪われる。自由が奪われるのは僕が一番嫌うことだ。自由が奪われるくらいなら一人がいい。"

        // // Detects the sentiment of the text.
        // sentiment, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
        //         Document: &languagepb.Document{
        //                 Source: &languagepb.Document_Content{
        //                         Content: text,
        //                 },
        //                 Type: languagepb.Document_PLAIN_TEXT,
        //         },
        //         EncodingType: languagepb.EncodingType_UTF8,
        // })

        sentiment, err := analyzeSentiment(ctx, client, text)

        if err != nil {
                log.Fatalf("Failed to analyze text: %v", err)
        }

        fmt.Println("--- Sentiment Analysis ---")
        fmt.Printf("Text: %v\n", text)
        fmt.Printf("Sentiment Score(-1<=score<=1): %v\n", sentiment.DocumentSentiment.Score)
        fmt.Printf("Magnitude: %v\n", sentiment.DocumentSentiment.Magnitude)
        // if sentiment.DocumentSentiment.Score >= 0 {
        //         fmt.Println("Sentiment: positive")
        // } else {
        //         fmt.Println("Sentiment: negative")
        // }

        entitySentiment, err := analyzeEntitySentiment(ctx, client, text)

        if err != nil {
                log.Fatalf("Failed to analyze text: %v", err)
        }

        fmt.Println("--- Entity Sentiment Analysis ---")
        fmt.Println("Language: %v\n", entitySentiment.Language)
        // for _, entity := entitySentiment.Entities {
        //         fmt.Printf("Text: %v\n", text)
        //         fmt.Printf("Entity: %v\n", entity.name)
        // }
        

}

func analyzeSentiment(ctx context.Context, client *language.Client, text string) (*languagepb.AnalyzeSentimentResponse, error) {
        return client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
                Document: &languagepb.Document{
                        Source: &languagepb.Document_Content{
                                Content: text,
                        },
                        Type: languagepb.Document_PLAIN_TEXT,
                },
        })
}

func analyzeEntitySentiment(ctx context.Context, client *language.Client, text string) (*languagepb.AnalyzeEntitySentimentResponse, error) {
        return client.AnalyzeEntitySentiment(ctx, &languagepb.AnalyzeEntitySentimentRequest{
                Document: &languagepb.Document{
                        Source: &languagepb.Document_Content{
                                Content: text,
                        },
                        Type: languagepb.Document_PLAIN_TEXT,
                },
        })
}
