import { useState } from 'react'
import { Article, Results, SearchResults } from './types/News'
import './App.css'
import axios from 'axios'

function App() {
  const [searchInput, setSearchInput] = useState<string>("")
  const [searchResults, setSearchResults] = useState<SearchResults>()

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchInput(e.target.value)
  }
  
  const handleOnClick = (e: React.MouseEvent<HTMLButtonElement>) => {
    const url = "http://localhost:8080/search?topic=" + searchInput 
    axios.get(url).then((response) => {  
      setSearchResults(processSearchResults(response.data))
      
    })  
    e.preventDefault()
  }

  const processSearchResults = (news: SearchResults) => {
    const processedArticles: Article[] = [];
    news.Results.articles.forEach((article) => {
      if (article.title !== "[Removed]") {
        processedArticles.push(article)
      }
    }) 

    const updatedResults = {} as Results
    updatedResults.status = news.Results.status
    updatedResults.articles = processedArticles
    updatedResults.totalResults = news.Results.totalResults

    const updatedSearchResults = {} as SearchResults
    updatedSearchResults.NextPage = news.NextPage
    updatedSearchResults.Query = news.Query
    updatedSearchResults.Results = updatedResults
    updatedSearchResults.TotalPages = news.TotalPages

    return updatedSearchResults
  }

  return (
    <>
      <div className="search-container">
        <form id="search-form" action="search" method="GET">
          <input 
            id="search-bar"
            name="topic"
            type="text"
            placeholder="Search the news regarding..."
            value={searchInput}
            onChange={handleChange}
          />
          <button id="search-button" type="submit" onClick={handleOnClick}> Search </button>
        </form>
      </div>
      <ul className="search-results"> 
        {searchResults?.Results.articles && searchResults.Results.articles.map((article, idx) => (
          <ul className="news-articles" key={`article-${idx}`}>
            <div>
              <a target="_blank" rel="noreferrer noopener" key={`url-${idx}`} href={article.url}>
                <h3 key={`title-${idx}`}>{article.title}</h3>
              </a>
              <p key={`description-${idx}`}>{article.description}</p>            
            </div>
            <img key={`image-${idx}`} src={article.urlToImage}/>
          </ul>
        ))}
      </ul>
    </>
  )
}

export default App
