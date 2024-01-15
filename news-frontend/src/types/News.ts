export declare interface Source {
  id:   string;
	name: string;
}

export declare interface Article {
  source:				Source;
	author:     	string;  
	title:       	string;   
	description: 	string;  
	url:         	string; 
	urkToImage:  	string;   
	publishedAt: 	string;
	content:     	string;
}

export declare interface Results {
  status:       string;  
	totalResults: number;      
	articles:     Article[];
}

export declare interface SearchResults {
	Query: 			string;
	NextPage: 	number;
	TotalPages: number;
	Results: 		Results;
}