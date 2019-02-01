components {
  id: "Road"
  component: "/Background/Road.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/Background/background.atlas\"\n"
  "default_animation: \"Asphalt\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: -148.0
    y: -0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite1"
  type: "sprite"
  data: "tile_set: \"/Background/background.atlas\"\n"
  "default_animation: \"Asphalt\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 108.0
    y: -0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.99999607
    w: 0.0028048193
  }
}
