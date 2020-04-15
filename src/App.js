import React from "react"
import axios from "axios"

class App extends React.Component {
	constructor() {
		super()
		this.state = { searchtext: "", countries: [] }
	}

	Country = (props) => {
		return (
			<div className="country center">
				<p className="country-text">{props.name}</p>
				<p>Cases: {props.cases}</p>
				<p>Deaths: {props.deaths}</p>
				<p>Recovered: {props.recovered}</p>
			</div>
		)
	}

	search = (text) => {
		let lowerText = text.toLowerCase()
		let modifiedText = lowerText.replace(" ", "-")
		if (text != "" || text != " ") {
			axios.get("http://localhost:5000/api/" + modifiedText)
				.then(res => {
					let countriesCopy = []
					countriesCopy.push([res.data.Cases, res.data.Deaths, res.data.Recovered, this.state.searchtext])
					this.setState({ countries: countriesCopy })
				})
		}
	}

	handleChange = (event) => {
		let { value } = event.target
		this.setState({ searchtext: value })
	}

	handleClick = () => {
		this.search(this.state.searchtext)
	}

	render() {
		const countries = this.state.countries.map(country => <this.Country cases={country[0]} deaths={country[1]} recovered={country[2]} name={country[3]} />)
		return (
			<div>
				<header className="center">
					<p>Covid Tracker</p>
				</header>
				<div className="center">
					<p className="info">Type the name of country you want to get stats for</p>
					<input type="text" name="searchtext" onChange={this.handleChange} className="form-control searchbox" />
					<button type="button" className="btn btn-outline-light" onClick={this.handleClick}>Search</button>
					{countries}
				</div>
			</div>
		);
	}
}

export default App