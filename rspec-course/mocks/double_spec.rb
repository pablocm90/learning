RSpec.describe 'a random double' do
  it 'only allows defined methods to be invoked' do
    stuntman = double(
      'Mr. Danger',
      {
        fall_of_ladder: 'AUCH!',
        light_on_fire: true
      }
    )
    expect(stuntman.fall_of_ladder).to eq('AUCH!')
    expect(stuntman.light_on_fire).to eq(true)
  end
end

RSpec.describe 'another random double' do
  it 'only allows defined methods to be invoked' do
    stuntman = double('Mr. Danger')

    allow(stuntman).to receive(:fall_of_ladder).and_return('AUCH!')
    expect(stuntman.fall_of_ladder).to eq('AUCH!')
    # expect(stuntman.light_on_fire).to eq(true)
  end
end

RSpec.describe 'a third random double' do
  it 'only allows defined methods to be invoked' do
    stuntman = double('Mr. Danger')

    allow(stuntman).to receive_messages({
                                          fall_of_ladder: 'AUCH!',
                                          light_on_fire: true
                                        })
    expect(stuntman.fall_of_ladder).to eq('AUCH!')
    expect(stuntman.light_on_fire).to eq(true)
  end
end

RSpec.describe 'doubles' do
  it "tests the student's knowledge of the course's content" do
    db = double(
      'Database Connection', 
      {
        connect: true,
        disconnect: 'Goodbye'
      }
    )
    
    expect(db.connect).to eq(true)
    expect(db.disconnect).to eq('Goodbye')
 
  end

  it "tests the student's knowledge of the course's content" do
    fs = double('File System')

    allow(fs).to receive_messages({
      read: "Romeo and Juliet",
      write: false
    })
 
  end
end
