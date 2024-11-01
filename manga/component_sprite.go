package manga

import (
	"github.com/veandco/go-sdl2/sdl"
)

const SpriteComponentType = "SpriteComponent"

type Animation struct {
	index     int32
	numFrames int32
	speed     int32
	isFixed   bool
}

type SpriteComponent struct {
	Entity        *Entity
	transform     *TransformComponent
	texture       *Texture
	src           *sdl.Rect
	dst           *sdl.Rect
	flip          sdl.RendererFlip
	animations    map[string]*Animation
	currAnimation *Animation
}

func makeSpriteComponent(entity *Entity, texture string) *SpriteComponent {
	t := Engine.AssetManager.GetTexture(texture)

	return &SpriteComponent{
		Entity:     entity,
		transform:  nil,
		texture:    t,
		src:        nil,
		dst:        nil,
		flip:       sdl.FLIP_NONE,
		animations: make(map[string]*Animation),
	}
}

func (s *SpriteComponent) SetTexture(id string) {
	s.texture = Engine.AssetManager.GetTexture(id)
}

func (s *SpriteComponent) Initialize() {
	s.transform = GetTransformComponent(s.Entity)
	s.src = &sdl.Rect{X: 0, Y: 0, W: 0, H: 0}
	s.dst = &sdl.Rect{X: 0, Y: 0, W: 0, H: 0}
}

func (s *SpriteComponent) Update() {
	s.src.W = s.transform.size.X
	s.src.H = s.transform.size.Y

	if s.currAnimation != nil {
		s.src.X = s.src.W * ((int32(sdl.GetTicks64()) / s.currAnimation.speed) % s.currAnimation.numFrames)
		s.src.Y = s.currAnimation.index * s.src.H
	}

	s.dst.X = int32(s.transform.position.X)
	s.dst.Y = int32(s.transform.position.Y)
	s.dst.W = s.transform.size.X * s.transform.scale
	s.dst.H = s.transform.size.Y * s.transform.scale
}

func (s *SpriteComponent) Render() {
	Engine.TextureManager.Draw(s.texture, s.src, s.dst, s.flip)
}

func (s *SpriteComponent) AddAnimation(name string, index int32, numFrames int32, speed int32, isFixed bool) {
	s.animations[name] = &Animation{
		index:     index,
		numFrames: numFrames,
		speed:     speed,
		isFixed:   isFixed,
	}
}

func (s *SpriteComponent) PlayAnimation(name string) {
	if animation, ok := s.animations[name]; ok {
		s.currAnimation = animation
	}
}

func GetSpriteComponent(entity *Entity) *SpriteComponent {
	component := entity.GetComponent(SpriteComponentType)

	if component == nil {
		return nil
	}

	return component.(*SpriteComponent)
}
