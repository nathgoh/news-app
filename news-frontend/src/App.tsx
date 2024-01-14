import { useEffect, useState } from 'react'
import { SearchResults } from './types/News'
import './App.css'
import axios from 'axios'

function App() {
  const [searchInput, setSearchInput] = useState<string>("")
  const [searchResults, setSearchResults] = useState<SearchResults>()

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchInput(e.target.value)
  }
  
  const handleOnSubmit = () => {
    const url = "http://localhost:8080/search?topic=" + searchInput 
    axios.get(url).then((response) => {
      setSearchResults(response.data as SearchResults)
    })
  }

  return (
    <>
      <div className="search-container">
        <form id="search-form" onSubmit={handleOnSubmit} action="search" method="GET">
          <input 
            id="search-bar"
            name="topic"
            type="text"
            placeholder="Search the news regarding..."
            value={searchInput}
            onChange={handleChange}
          />
          <button id="search-button" type="submit"> Search </button>
        </form>
      </div>
      <div className="search-results"> 
        {searchResults}
      </div>
    </>
  )
}

export default App
