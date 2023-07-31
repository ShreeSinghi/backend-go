const express = require("express");
const router = express.Router();
const { authenticate, getDataAdmin,  getDataUser} = require('../middleware');
const db = require("../database");

router.post('/request-checkout', authenticate, (req, res) => {
    const bookId = req.body.bookId
    const userId = req.body.userId
  
    db.query(`SELECT * FROM books WHERE id=${db.escape(bookId)}`, async (err, results) => {
      if (err) throw err
      if (results.length === 0) return res.send(await getDataUser(userId, 'Book does not exist'))
      if (results[0].quantity === 0) return res.send(await getDataUser(userId, 'Book is out of stock'))
  
      db.query(`SELECT * FROM requests WHERE bookId=${db.escape(bookId)} AND userId=${userId} AND state='outrequested'`,
        async (err, results) => {
          if (err) throw err
          if (results.length > 0) return res.send(await getDataUser(userId, 'You have already requested this book'))
  
          db.query(`INSERT INTO requests (bookId, userId, state) VALUES (${db.escape(bookId)}, ${userId}, 'outrequested')`,
          async (err, results) => {
              if (err) throw err
              res.send(await getDataUser(userId, 'Checkout request submitted'))
            }
          )
  
          db.query(`UPDATE books SET quantity = quantity - 1 WHERE id = ${db.escape(bookId)}`, (err, results) => {
            if (err) throw err
          })
        }
      )
    })
  })
  
  router.post('/return-book', authenticate, (req, res) => {
    const bookId = req.body.bookId
    const userId = req.body.userId
  
    db.query(`SELECT * FROM requests WHERE bookId=${db.escape(bookId)} AND userId=${userId} AND state='owned'`, (err, results) => {
        if (err) throw err
  
        console.log(results)
  
        if (results.length === 0) {
          console.log('boook does not exist or is not owned by the user')
          return res.redirect('home')
  
        }
    
        db.query( `UPDATE requests SET state = 'inrequested' WHERE id=${results[0].id}`, (err) => {
            if (err) throw err
            console.log('bbook return requested')
            return res.redirect('home')
          }
        )
      }
    )
  })
  

  
router.post('/request-admin', authenticate, (req, res) => {
    const userId = req.body.userId
  
    if (req.body.admin) return res.status(403).send({ msg: 'User is already an admin' })
  
    db.query(`UPDATE users SET requested = true WHERE id = ${db.escape(userId)}`,
      (err, results) => {
        if (err) throw err
        console.log('hii')
        res.send(getDataUser(userId, ''))
      }
    )
  })
  
  module.exports = router;