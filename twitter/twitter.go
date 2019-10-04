package twitter

import (
    "fmt"
    "log"
    "net/url"
    "os"
    "os/signal"

    "github.com/ChimeraCoder/anaconda"
    "github.com/joho/godotenv"
    "time"
    // "sort"
)

type tagNum struct{
    text string
    count int
}
type tagNums []tagNum
type tagImg struct{
    text string
    img []string
}
type tagStr struct{
    text string
    count int
    images []string
    point int
}

func loadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func getTwitterApi() *anaconda.TwitterApi {
    anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
    anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
    return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}

func Sort(mstr []tagStr) []tagStr { 
    so := 1
    if len(mstr) == 0 || len(mstr) == 1 {
    	return mstr
    }
    for {
        if mstr[so - 1].point < mstr[so].point {
            mstr[so - 1], mstr[so] = mstr[so], mstr[so -1]
            so = 0
        }
        so = so + 1
        if so == len(mstr) {
            break
        }
    }
    return mstr
}

func arrayToStruct(array []tagImg) []tagStr {
    m := map[string]int{}
    var newTagStr tagStr
    var m_struct []tagStr
    var m_structt []tagStr
    for _, arr := range array {
        _, ok := m[arr.text]
        if ok {
            m[arr.text] = m[arr.text] + 1
        }else{
            m[arr.text] = 1
        }
    }
    for tex, cou := range m {
        newTagStr.count = cou
        newTagStr.text = tex
        m_struct = append(m_struct, newTagStr)
    }
    for _, tagim := range m_struct {
        for _, artagim := range array {
            if tagim.text == artagim.text {
                // tagim.images = append(tagim.images, artagim.img)
                for _, img := range artagim.img {
                    tagim.images = append(tagim.images, img)
                }
            }
        }
        tagim.point = tagim.count * len(tagim.images)
        m_structt = append(m_structt, tagim)
    }
    // m_structt = compress(m_structt)
    return Sort(m_structt)
}

func GetTweet() {
    loadEnv()

    api := getTwitterApi()

    v := url.Values{}
    v.Set("count", "100")
    v.Set("lang", "ja")
    var tags_form []tagImg
    var tag_form tagImg
    var images []string
    i := 1
        for {searchResult, _ := api.GetSearch("%23", v)
            for _, tweet := range searchResult.Statuses {
                images = nil
                medias := tweet.Entities.Media
                for _, media := range medias {
                    images = append(images, media.Media_url)
                }
                tags := tweet.Entities.Hashtags
                for _, tag := range tags {
                    if tag.Text != ""  {
                        tag_form.text = tag.Text
                        for _, img := range images{
                            tag_form.img = append(tag_form.img, img)
                        }
                        tags_form = append(tags_form, tag_form)
                        tag_form.img = nil
                    }
                }
            }
            fmt.Println(i)
            if i % 40 == 0 {
            	if len(arrayToStruct(tags_form)) > 2{
                	for k := 0; k < 10; k++{
                    	fmt.Println(arrayToStruct(tags_form)[k])
                	}
                	tags_form = nil
            	}else{
        		fmt.Println("十分な数のツイートを取得できませんでした。")
        		}
        	}
            i = i + 1
            time.Sleep(5 * time.Second)
        }
    // シグナル用のチャネル定義
    quit := make(chan os.Signal)
    // 受け取るシグナルを設定
    signal.Notify(quit, os.Interrupt)
    <-quit // ここでシグナルを受け取るまで以降の処理はされない

    // シグナルを受け取った後にしたい処理を書く
    fmt.Println("終了しました。")
}
