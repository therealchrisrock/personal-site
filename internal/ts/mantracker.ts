import * as d3 from "d3";

export function initD3Visualization() {
    const data = [10, 20, 30, 40, 50];
    const width = 400, height = 200;

    const svg = d3.select("#d3-container")
        .append("svg")
        .attr("width", width)
        .attr("height", height);

    svg.selectAll("circle")
        .data(data)
        .enter()
        .append("circle")
        .attr("cx", (d, i) => i * 80 + 40)
        .attr("cy", height / 2)
        .attr("r", d => d / 2)
        .style("fill", "steelblue");
}