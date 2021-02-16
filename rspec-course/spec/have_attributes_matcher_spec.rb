class ProfessionalWrestler
  attr_reader :name, :finishing_move
  def initialize(name, finishing_move)
    @name = name
    @finishing_move = finishing_move
  end
end

RSpec.describe 'have attributes matcher' do 
  describe ProfessionalWrestler.new('Pablo', 'Tortilla') do 
    it 'checks for object attributes and proper values' do 
      expect(subject).to have_attributes(name: 'Pablo')
      expect(subject).to have_attributes(name: 'Pablo', finishing_move: 'Tortilla')
    end
    it { is_expected.to have_attributes(name: 'Pablo') }
  end
end