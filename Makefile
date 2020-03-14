.PHONY: clean run play docker/build
DOCKER_IMAGE := jewel12/magenta

BUNDLE_FILE := data/chord_pitches_improv.mag
OUTPUT_DIR := data/generated
SOUND_FONT := data/TimGM6mb.sf2

run:
	go run server.go

docker/build:
	docker build . -t ${DOCKER_IMAGE}

play: clean docker/build ${BUNDLE_FILE} ${OUTPUT_DIR} ${SOUND_FONT}
	docker run -v $(shell pwd)/data:/data --rm ${DOCKER_IMAGE} improv_rnn_generate \
		--config=chord_pitches_improv \
		--bundle_file=/${BUNDLE_FILE} \
		--output_dir=/${OUTPUT_DIR} \
		--num_outputs=1 \
		--primer_melody=${MELODY} \
		--backing_chords=${CHORD} \
		--render_chords
	fluidsynth -i ${SOUND_FONT} ${OUTPUT_DIR}/*.mid

${BUNDLE_FILE}: data
	curl http://download.magenta.tensorflow.org/models/chord_pitches_improv.mag > $@

${OUTPUT_DIR}:
	-mkdir -p ${OUTPUT_DIR}

data:
	-mkdir $@

${SOUND_FONT}:
	curl http://kujirahand.com/download/2014/TimGM6mb.sf2 > $@

clean:
	-rm -r data/generated
