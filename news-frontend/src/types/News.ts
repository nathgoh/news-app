export declare interface Source {
  ID:   string;
	Name: string;
}

export declare interface Article {
  Source: Source;
	Author:      string;  
	Title:      string;   
	Description: string;  
	URL:         string; 
	URLToImage:  string;   
	PublishedAt: string;
	Content:     string;
}

export declare interface SearchResults {
  Status:       string;  
	TotalResults: number;      
	Articles:     Article[];
}