package web

const MainJs = `
// Chart.js axis labels don't support natively word wrapping, so here's a workaround:
const wrap = (str, limit) => {
  const words = str.split(' ');
  let aux = [];
  let concat = [];

  for (let i = 0; i < words.length; i++) {
    concat.push(words[i]);
    let join = concat.join(' ');
    if (join.length > limit) {
      aux.push(join);
      concat = [];
    }
  }

  if (concat.length) {
    aux.push(concat.join(' ').trim());
  }

  return aux;
}

const createCanvas = id => {
  let newDiv = document.createElement('div');
  newDiv.className = 'messagesWrapper';
  let newCanvas = document.createElement('canvas');
  newCanvas.setAttribute('id', id);
  newDiv.appendChild(newCanvas);
  document.getElementById('divWorkflowMessages').appendChild(newDiv);
}

const jobsSuccessRates = (id, labels, datasets) => {
  var ctx = document.getElementById(id).getContext('2d');
  new Chart(ctx, {
      type: 'horizontalBar',
      data: {
        labels: labels,
        datasets: [{
            data: datasets,
            barThickness: 25
        }]
      },
      options: {
          title: {
            display: false,
          },
          scales: {
            xAxes: [{
              ticks: {
                beginAtZero: true
              }
            }]
          },
          legend: {
            display: false
          },
          layout: {
            padding: 0
          }
      }
  });
}

const workflowMessages = (workflow, labels, datasets) => {
  var ctx = document.getElementById(workflow.Id).getContext('2d');
  new Chart(ctx, {
      type: 'horizontalBar',
      data: {
          labels: labels,
          datasets: [{
              data: datasets,
              barThickness: 25
          }]
      },
      options: {
          title: {
            display: true,
            fontSize: 20,
            text: workflow.Name
          },
          scales: {
            yAxes: [{
              ticks: {
                beginAtZero: true,
                callback: function(label) {
                  return wrap(label, 40);
                }
              }
            }]
          },
          legend: {
            display: false
          },
          layout: {
            padding: 0
          }
      }
  });
}
`