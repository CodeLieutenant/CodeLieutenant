import { tsParticles } from 'tsparticles';

const particlesSnowEffect = async (snowItem: string) => {
  await tsParticles.load(snowItem, {
    particles: {
      number: {
        value: 100,
        density: {
          enable: true,
          value_area: 800,
        },
      },
      color: {
        value: '#ffffff',
      },
      shape: {
        type: 'image',
        stroke: {
          width: 3,
          color: '#fff',
        },
        polygon: {
          nb_sides: 5,
        },
        image: {
          src: '/public/img/snow_flake.png',
          width: 100,
          height: 100,
        },
      },
      opacity: {
        value: 0.7,
        random: true,
        anim: {
          enable: true,
          speed: 1,
          opacity_min: 0.1,
          sync: false,
        },
      },
      size: {
        value: 5,
        random: true,
        anim: {
          enable: false,
          speed: 20,
          size_min: 0.1,
          sync: false,
        },
      },
      line_linked: {
        enable: false,
        distance: 50,
        color: '#ffffff',
        opacity: 0.6,
        width: 1,
      },
      move: {
        enable: true,
        speed: 5,
        direction: 'bottom',
        random: true,
        straight: false,
        out_mode: 'out',
        bounce: false,
        attract: {
          enable: true,
          rotateX: 300,
          rotateY: 1200,
        },
      },
    },
    retina_detect: true,
  });
};

export { particlesSnowEffect };
