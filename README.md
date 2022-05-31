# Animal Soundboard

📢 A simple soundboard of animal sounds

## Data

[Navigate to seaworld.org](https://seaworld.org/animals/sounds/)

```js
$$('.section-content-item__title-description-wrapper').map(node => {
  return [node.querySelector('h3').textContent, node.querySelector('audio source').src]
}) 
```
