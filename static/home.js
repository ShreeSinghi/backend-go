console.log("gay1")

document.addEventListener('DOMContentLoaded', () => {
  const requestForm = document.getElementById('checkout-form');
  const requestModal = new bootstrap.Modal(document.getElementById('requestModal'));
  const requestStatus = document.getElementById('requestStatus');
  console.log("gay2")

  requestForm.addEventListener('submit', (event) => {
    event.preventDefault();
    console.log("mewoowwowo")

    const bookId = document.getElementById('requestBook').value;

    axios.post('/request-checkout' + bookId.toString(), { bookId }, {
      headers: {
        'Content-Type': 'application/json'
      }
    })
    .then(response => {
      const data = response.data;
      console.log(data.checkoutStatus);
      requestStatus.textContent = data.checkoutStatus;

      if (data.checkoutStatus.includes('submitted')) {
        requestStatus.style.color = 'green';
      } else {
        requestStatus.style.color = 'red';
      }

      requestModal.show();
    })
    .catch(error => {
      console.error('Error:', error);
    });
  });

  const requestAdminBtn = document.getElementById('requestAdminBtn');

  requestAdminBtn.addEventListener('click', (event) => {
    event.preventDefault();

    axios.post('/request-admin')
    .then(response => {
      location.reload();
    })
    .catch(error => {
      console.error('Error:', error);
    });
  });

  const returnBtn = document.getElementById('returnBtn');

  returnBtn.addEventListener('click', (event) => {
    event.preventDefault();
    const bookId = document.getElementById('returnBook').value;

    axios.post('/return-book', { bookId })
    .then(response => {
      location.reload();
    })
    .catch(error => {
      console.error('Error:', error);
    });
  });

  const closeModalBtn = document.getElementById('closeModalBtn');

  closeModalBtn.addEventListener('click', () => {
    location.reload();
  });
});